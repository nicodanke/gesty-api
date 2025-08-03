package gapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nicodanke/gesty-api/services/employee-service/clients"
	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/employee-service/sse"
	"github.com/nicodanke/gesty-api/services/employee-service/utils"
	es "github.com/nicodanke/gesty-api/shared/proto/employee-service"
	"github.com/nicodanke/gesty-api/shared/token"
)

type Server struct {
	es.UnimplementedEmployeeServiceServer
	config         utils.Config
	store          db.Store
	tokenMaker     token.Maker
	notifier       sse.Notifier
	deepfaceClient *clients.DeepfaceClient
}

func NewServer(config utils.Config, store db.Store, notifier sse.Notifier) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	deepfaceClient := &clients.DeepfaceClient{
		BaseURL: config.DeepfaceServiceAddress,
		Client:  &http.Client{Timeout: time.Second * 50},
	}

	server := &Server{store: store, tokenMaker: tokenMaker, config: config, notifier: notifier, deepfaceClient: deepfaceClient}

	return server, nil
}
