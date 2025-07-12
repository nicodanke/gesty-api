package gapi

import (
	"context"
	"fmt"
	"strings"

	"github.com/nicodanke/gesty-api/shared/token"
	"google.golang.org/grpc/metadata"
)

const (
	authorizationHeader     = "authorization"
	authorizationTypeBearer = "bearer"
)

func (server *Server) authenticateUser(ctx context.Context) (*token.Payload, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	values := md.Get(authorizationHeader)
	if len(values) == 0 {
		return nil, fmt.Errorf("missing authorization header")
	}

	authHeader := values[0]
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationTypeBearer {
		return nil, fmt.Errorf("unsupported authorization type %s", authType)
	}
	// The JWT token is the second field in the authorization header
	tokenString := fields[1]

	payload, err := server.tokenMaker.VerifyToken(tokenString)
	if err != nil {
		return nil, fmt.Errorf("inavalid access token: %s", err)
	}

	return payload, nil
}