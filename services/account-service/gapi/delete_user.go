package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/gesty-api/services/account-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/account-service/requests/user"
	"github.com/nicodanke/gesty-api/services/account-service/sse"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	sse_delete_user = "delete-user"
)

func (server *Server) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*emptypb.Empty, error) {
	log.Info().Str("method", "DeleteUser").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing DeleteUser request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAU", "DU"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAU or DU"))
	}

	if req.GetId() == authPayload.UserID {
		return nil, unprocessableError(ACTION_NOT_ALLOWED, "User cannot auto delete itself")
	}

	arg := db.DeleteUserParams{
		AccountID: authPayload.AccountID,
		ID:        req.GetId(),
	}

	err = server.store.DeleteUser(ctx, arg)
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to delete user with id:", req.GetId(), err))
	}

	// Notify delete user
	var data = map[string]any{}
	data["id"] = req.GetId()
	server.notifier.BoadcastMessageToAccount(sse.NewEventMessage(sse_delete_user, data), authPayload.AccountID, &authPayload.UserID)

	return &emptypb.Empty{}, nil
}