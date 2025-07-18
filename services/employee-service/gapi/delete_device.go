package gapi

import (
	"context"
	"fmt"
	"strconv"

	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/employee-service/sse"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/device"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	sse_delete_device = "delete-device"
)

func (server *Server) DeleteDevice(ctx context.Context, req *device.DeleteDeviceRequest) (*emptypb.Empty, error) {
	log.Info().Str("method", "DeleteDevice").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing DeleteDevice request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAD", "LD"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAD or LD"))
	}

	argDelete := db.DeleteDeviceTxParams{
		AccountID: authPayload.AccountID,
		ID:        req.GetId(),
	}

	err = server.store.DeleteDeviceTx(ctx, argDelete)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.ForeignKeyViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_FK, fmt.Sprintln("Failed to delete device due to foreign key constraint violation"), constraintName)
		}

		return nil, internalError(fmt.Sprintln("Failed to delete device:", err))
	}

	// Notify delete device
	var data = map[string]any{}
	data["id"] = strconv.FormatInt(req.GetId(), 10)
	server.notifier.BoadcastMessageToAccount(sse.NewEventMessage(sse_delete_device, data), authPayload.AccountID, &authPayload.UserID)

	return &emptypb.Empty{}, nil
}
