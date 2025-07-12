package main

import (
	"context"
	"net"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	_ "github.com/nicodanke/gesty-api/services/employee-service/doc/statik"
	"github.com/nicodanke/gesty-api/services/employee-service/gapi"
	gw "github.com/nicodanke/gesty-api/services/employee-service/gw/employee-service"
	"github.com/nicodanke/gesty-api/services/employee-service/sse"
	"github.com/nicodanke/gesty-api/services/employee-service/utils"
	ac "github.com/nicodanke/gesty-api/shared/proto/employee-service"
	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := utils.LoadConfig()
	if err != nil {
		log.Error().Err(err).Msg("Cannot load configuration")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Error().Err(err).Msg("Cannot connect to db")

	}

	runDBMigrations(config.MigrationUrl, config.DBSource)

	store := db.NewStore(connPool)

	// Creates HandlerEvent to send events through HTTP Server Sent Events (SSE)
	handlerEvent := sse.NewHandlerEvent()

	go runGRPCGatewayServer(config, store, handlerEvent)
	go runServerSentEvents(config, handlerEvent)
	runGRPCServer(config, store, handlerEvent)
}

func runDBMigrations(migrationUrl string, dbSource string) {
	migration, err := migrate.New(migrationUrl, dbSource)
	if err != nil {
		log.Error().Err(err).Msg("Cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Error().Err(err).Msg("Failed to run migrate up")
	}

	log.Info().Msg("DB migrations runned successfully")
}

func runGRPCServer(config utils.Config, store db.Store, notifier sse.Notifier) {
	server, err := gapi.NewServer(config, store, notifier)
	if err != nil {
		log.Error().Err(err).Msg("Cannot create server")
	}

	grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	ac.RegisterEmployeeServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Error().Err(err).Msg("Cannot create listener")
	}

	log.Info().Msgf("gRPC server started at: %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Error().Err(err).Msg("Cannot start gRPC server")
	}
}

func runGRPCGatewayServer(config utils.Config, store db.Store, notifier sse.Notifier) {
	server, err := gapi.NewServer(config, store, notifier)
	if err != nil {
		log.Error().Err(err).Msg("Cannot create server")
	}

	grpcMux := runtime.NewServeMux()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = gw.RegisterEmployeeServiceHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Error().Err(err).Msg("Cannot register handler server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	statikFS, err := fs.New()
	if err != nil {
		log.Error().Err(err).Msg("Cannot create static file system")
	}

	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
	mux.Handle("/swagger/", swaggerHandler)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Error().Err(err).Msg("Cannor create listener")
	}

	log.Info().Msgf("HTTP Gateway Server started at: %s", listener.Addr().String())

	muxWithCors := cors.New(cors.Options{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"ACCEPT", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler(mux)

	handler := gapi.HttpLogger(muxWithCors)
	err = http.Serve(listener, handler)
	if err != nil {
		log.Error().Err(err).Msg("Cannor start HTTP Gateway Server")
	}
}

func runServerSentEvents(config utils.Config, handlerEvent *sse.HandlerEvent) {
	server, err := sse.NewServer(config, handlerEvent)
	if err != nil {
		log.Error().Err(err).Msg("Cannot create HTTP SSE server")
	}

	err = server.Start(config.SSEAddress)
	if err != nil {
		log.Error().Err(err).Msg("Cannot start HTTP SSE server")
	}

	log.Info().Msg("HTTP SSE server started")
}
