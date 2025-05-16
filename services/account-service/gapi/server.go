package gapi

import (
	"fmt"

	db "github.com/nicodanke/gesty-api/services/account-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/account-service"
	"github.com/nicodanke/gesty-api/services/account-service/sse"
	"github.com/nicodanke/gesty-api/shared/token"
	"github.com/nicodanke/gesty-api/services/account-service/utils"
)

type Server struct {
	pb.UnimplementedInventAppV1Server
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
	notifier   sse.Notifier
}

func NewServer(config utils.Config, store db.Store, notifier sse.Notifier) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{store: store, tokenMaker: tokenMaker, config: config, notifier: notifier}

	return server, nil
}