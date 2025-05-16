package main

import (
	"context"
	// "net"
	// "net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/nicodanke/gesty-api/services/account-service/utils"
	db "github.com/nicodanke/gesty-api/services/account-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/account-service/sse"
	"github.com/nicodanke/gesty-api/services/account-service/gapi"
	"github.com/nicodanke/gesty-api/shared/proto/account-service"
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

	db.NewStore(connPool)

	// Creates HandlerEvent to send events through HTTP Server Sent Events (SSE)
	handlerEvent := sse.NewHandlerEvent()

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
	pb.RegisterInventAppV1Server(grpcServer, server)
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
