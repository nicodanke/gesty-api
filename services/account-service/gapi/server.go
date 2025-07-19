package gapi

import (
	"fmt"

	db "github.com/nicodanke/gesty-api/services/account-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/account-service/sse"
	"github.com/nicodanke/gesty-api/services/account-service/utils"
	as "github.com/nicodanke/gesty-api/shared/proto/account-service"
	es "github.com/nicodanke/gesty-api/shared/proto/employee-service"
	"github.com/nicodanke/gesty-api/shared/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	as.UnimplementedAccountServiceServer
	config         utils.Config
	store          db.Store
	tokenMaker     token.Maker
	notifier       sse.Notifier
	employeeClient es.EmployeeServiceClient
}

func NewServer(config utils.Config, store db.Store, notifier sse.Notifier) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	conn, err := grpc.NewClient(config.GRPCEmployeeServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("cannot dial employee service: %w", err)
	}
	// defer conn.Close()

	employeeClient := es.NewEmployeeServiceClient(conn)

	server := &Server{store: store, tokenMaker: tokenMaker, config: config, notifier: notifier, employeeClient: employeeClient}

	return server, nil
}
