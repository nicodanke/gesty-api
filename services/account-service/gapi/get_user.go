package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/gesty-api/services/account-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/account-service/requests/user"
)

func (server *Server) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAU", "LU"}})
	if !authorized {
		return nil, permissionDeniedError("FORBIDDEN", fmt.Sprintln("User not authorized"))
	}

	arg := db.GetUserParams{
		AccountID: authPayload.AccountID,
		ID:        req.GetId(),
	}

	result, err := server.store.GetUser(ctx, arg)
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to get users:", err))
	}

	rsp := &user.GetUserResponse{
		User: convertUser(result),
	}
	return rsp, nil
}
