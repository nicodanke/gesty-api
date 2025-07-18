package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/device"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) RefreshDeviceToken(ctx context.Context, req *device.RefreshDeviceTokenRequest) (*device.RefreshDeviceTokenResponse, error) {
	log.Info().Str("method", "RefreshDeviceToken").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing RefreshDeviceToken request")

	payload, err := server.tokenMaker.VerifyTokenDevice(string(req.RefreshToken))
	if err != nil {
		return nil, invalidArgumentError([]*errdetails.BadRequest_FieldViolation{fieldViolation("refresh_token", err)})
	}

	deviceID := payload.DeviceID
	accountID := payload.AccountID

	deviceModel, err := server.store.GetDevice(ctx, db.GetDeviceParams{AccountID: accountID, ID: deviceID})
	if err != nil {
		return nil, conflictError("", fmt.Sprintln("Device not found"), "id")
	}

	if !deviceModel.Enabled {
		return nil, conflictError("", fmt.Sprintln("Device is not enabled"), "id")
	}

	token, payload, err := server.tokenMaker.CreateTokenDevice(deviceID, accountID, server.config.DeviceAccessTokenDuration)
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to create device token", err))
	}

	newRefreshToken, newRefreshPayload, err := server.tokenMaker.CreateTokenDevice(deviceID, accountID, server.config.DeviceRefreshTokenDuration)
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to create refresh token", err))
	}

	actionIds := make([]int64, 0)
	for _, v := range deviceModel.ActionIds.([]interface{}) {
		actionIds = append(actionIds, v.(int64))
	}

	rsp := &device.RefreshDeviceTokenResponse{
		AccessToken:           token,
		RefreshToken:          newRefreshToken,
		AccessTokenExpiresAt:  timestamppb.New(payload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(newRefreshPayload.ExpiredAt),
		ActionIds:             actionIds,
	}

	return rsp, nil
}
