package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/gesty-api/services/account-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/account-service/requests/role"
	"github.com/nicodanke/gesty-api/services/account-service/sse"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	sse_delete_role = "delete-role"
)

func (server *Server) DeleteRole(ctx context.Context, req *role.DeleteRoleRequest) (*emptypb.Empty, error) {
	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	arg := db.DeleteRoleParams{
		AccountID: authPayload.AccountID,
		ID:        req.GetId(),
	}

	err = server.store.DeleteRole(ctx, arg)
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to delete role with id:", req.GetId(), err))
	}

	// Notify delete role
	var data = map[string]any{}
	data["id"] = req.GetId()
	server.notifier.BoadcastMessageToAccount(sse.NewEventMessage(sse_delete_role, data), authPayload.AccountID, &authPayload.UserID)

	return &emptypb.Empty{}, nil
}