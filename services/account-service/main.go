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
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

	// go runGRPCGatewayServer(config, store, handlerEvent)
	// go runServerSentEvents(config, handlerEvent)
	// runGRPCServer(config, store, handlerEvent)
}