package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/action"
	"github.com/rs/zerolog/log"
)

func (server *Server) GetAction(ctx context.Context, req *action.GetActionRequest) (*action.GetActionResponse, error) {
	log.Info().Str("method", "GetAction").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing GetAction request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	if req.GetId() != authPayload.UserID {
		authorized := server.authorizeUser(authPayload, [][]string{{"SAA", "LA"}})
		if !authorized {
			return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAA or LA"))
		}
	}

	arg := db.GetActionParams{
		AccountID: authPayload.AccountID,
		ID:        req.GetId(),
	}

	result, err := server.store.GetAction(ctx, arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return nil, notFoundError("", fmt.Sprintln("Action with id", req.GetId(), "not found"))
		}
		return nil, internalError(fmt.Sprintln("Failed to get action:", err))
	}

	rsp := &action.GetActionResponse{
		Action: convertAction(result),
	}
	return rsp, nil
}
