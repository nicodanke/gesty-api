package sse

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nicodanke/inventapp_v3/account-service/token"
	"github.com/nicodanke/inventapp_v3/account-service/utils"
)

type Server struct {
	config       utils.Config
	tokenMaker   token.Maker
	handlerEvent *HandlerEvent
	router       *gin.Engine
}

func NewServer(config utils.Config, handlerEvent *HandlerEvent) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{tokenMaker: tokenMaker, config: config, handlerEvent: handlerEvent}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// Auth Required
	authRoutes.GET("/events", func(ctx *gin.Context) {
		server.handlerEvent.Handler(ctx)
	})

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}