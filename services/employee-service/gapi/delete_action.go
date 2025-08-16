package gapi

import (
	"context"
	"fmt"
	"strconv"

	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/employee-service/sse"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/action"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	sse_delete_action = "delete-action"
)

func (server *Server) DeleteAction(ctx context.Context, req *action.DeleteActionRequest) (*emptypb.Empty, error) {
	log.Info().Str("method", "DeleteAction").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing DeleteAction request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAA", "DA"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAA or DA"))
	}

	arg := db.GetActionParams{
		AccountID: authPayload.AccountID,
		ID:        req.GetId(),
	}

	action, err := server.store.GetAction(ctx, arg)
	if err != nil {
		return nil, notFoundError(NOT_FOUND, fmt.Sprintln("Action not found:", err))
	}

	if !action.CanBeDeleted {
		return nil, unprocessableError(ACTION_NOT_DELETABLE, fmt.Sprintln("Action cannot be deleted"))
	}

	argDelete := db.DeleteActionParams{
		AccountID: authPayload.AccountID,
		ID:        req.GetId(),
	}

	err = server.store.DeleteAction(ctx, argDelete)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.ForeignKeyViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_FK, fmt.Sprintln("Failed to delete action due to foreign key constraint violation"), constraintName)
		}

		return nil, internalError(fmt.Sprintln("Failed to delete action:", err))
	}

	// Notify delete action
	var data = map[string]any{}
	data["id"] = strconv.FormatInt(req.GetId(), 10)
	server.notifier.BoadcastMessageToAccount(sse.NewEventMessage(sse_delete_action, data), authPayload.AccountID, &authPayload.UserID)

	return &emptypb.Empty{}, nil
}
