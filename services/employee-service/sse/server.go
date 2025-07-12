package sse

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nicodanke/gesty-api/services/employee-service/utils"
	"github.com/nicodanke/gesty-api/shared/token"
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

	// Add CORS middleware
	router.Use(corsMiddleware())

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

// corsMiddleware handles CORS headers and OPTIONS preflight requests
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Expose-Headers", "Content-Length")
		c.Header("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
