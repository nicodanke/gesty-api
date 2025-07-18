package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/device"
	"github.com/rs/zerolog/log"
)

func (server *Server) GetDevice(ctx context.Context, req *device.GetDeviceRequest) (*device.GetDeviceResponse, error) {
	log.Info().Str("method", "GetDevice").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing GetDevice request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	if req.GetId() != authPayload.UserID {
		authorized := server.authorizeUser(authPayload, [][]string{{"SAD", "LD"}})
		if !authorized {
			return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAD or LD"))
		}
	}

	arg := db.GetDeviceParams{
		AccountID: authPayload.AccountID,
		ID:        req.GetId(),
	}

	result, err := server.store.GetDevice(ctx, arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return nil, notFoundError("", fmt.Sprintln("Device with id", req.GetId(), "not found"))
		}
		return nil, internalError(fmt.Sprintln("Failed to get device:", err))
	}

	rsp := &device.GetDeviceResponse{
		Device: convertGetDeviceRow(result),
	}
	return rsp, nil
}
