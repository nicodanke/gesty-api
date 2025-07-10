package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/gesty-api/services/account-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/account-service/requests/user"
	"github.com/rs/zerolog/log"
)

func (server *Server) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	log.Info().Str("method", "GetUser").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing GetUser request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	if req.GetId() != authPayload.UserID {
		authorized := server.authorizeUser(authPayload, [][]string{{"SAU", "LU"}})
		if !authorized {
			return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAU or LU"))
		}
	}

	arg := db.GetUserParams{
		AccountID: authPayload.AccountID,
		ID:        req.GetId(),
	}

	result, err := server.store.GetUser(ctx, arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return nil, notFoundError("", fmt.Sprintln("User with id", req.GetId(), "not found"))
		}
		return nil, internalError(fmt.Sprintln("Failed to get users:", err))
	}

	rsp := &user.GetUserResponse{
		User: convertUser(result),
	}
	return rsp, nil
}
