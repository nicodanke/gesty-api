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

	// go runGRPCGatewayServer(config, store, handlerEvent)
	runServerSentEvents(config, handlerEvent)
	// runGRPCServer(config, store, handlerEvent)
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
