package gapi

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/device"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) ActivateDevice(ctx context.Context, req *device.ActivateDeviceRequest) (*device.ActivateDeviceResponse, error) {
	log.Info().Str("method", "ActivateDevice").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing ActivateDevice request")

	code, err := base64.StdEncoding.DecodeString(req.GetCode())
	if err != nil {
		return nil, invalidArgumentError([]*errdetails.BadRequest_FieldViolation{fieldViolation("code", err)})
	}

	codeParts := strings.Split(string(code), "|")

	if len(codeParts) != 3 {
		return nil, invalidArgumentError([]*errdetails.BadRequest_FieldViolation{fieldViolation("code", fmt.Errorf("invalid activation code"))})
	}

	accountID, err := strconv.ParseInt(codeParts[0], 10, 64)
	if err != nil {
		return nil, invalidArgumentError([]*errdetails.BadRequest_FieldViolation{fieldViolation("code", err)})
	}

	deviceID, err := strconv.ParseInt(codeParts[1], 10, 64)
	if err != nil {
		return nil, invalidArgumentError([]*errdetails.BadRequest_FieldViolation{fieldViolation("code", err)})
	}

	if req.GetId() != deviceID {
		return nil, invalidArgumentError([]*errdetails.BadRequest_FieldViolation{fieldViolation("id", fmt.Errorf("id does not match activation code"))})
	}

	deviceModel, err := server.store.GetDevice(ctx, db.GetDeviceParams{AccountID: accountID, ID: deviceID})
	if err != nil {
		return nil, conflictError("", fmt.Sprintln("Device not found"), "id")
	}

	if !deviceModel.Enabled {
		return nil, conflictError("", fmt.Sprintln("Device is not enabled"), "id")
	}

	if deviceModel.ActivationCode.String != codeParts[2] {
		return nil, conflictError("", fmt.Sprintln("Invalid activation code"), "id")
	}

	if deviceModel.ActivationCodeExpiresAt.Before(time.Now()) {
		return nil, conflictError("", fmt.Sprintln("Activation code has expired"), "id")
	}

	token, payload, err := server.tokenMaker.CreateTokenDevice(deviceID, accountID, server.config.DeviceAccessTokenDuration)
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to create device token", err))
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateTokenDevice(deviceID, accountID, server.config.DeviceRefreshTokenDuration)
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to create refresh token", err))
	}

	_, err = server.store.UpdateDevice(ctx, db.UpdateDeviceParams{
		AccountID:               accountID,
		ID:                      deviceID,
		Active:                  pgtype.Bool{Bool: true, Valid: true},
		ActivationCodeExpiresAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt:               pgtype.Timestamptz{Time: time.Now(), Valid: true},
		DeviceName:              pgtype.Text{String: req.GetDeviceName(), Valid: req.DeviceName != ""},
		DeviceModel:             pgtype.Text{String: req.GetDeviceModel(), Valid: req.DeviceModel != ""},
		DeviceBrand:             pgtype.Text{String: req.GetDeviceBrand(), Valid: req.DeviceBrand != ""},
		DeviceSerialNumber:      pgtype.Text{String: req.GetDeviceSerialNumber(), Valid: req.DeviceSerialNumber != ""},
		DeviceOs:                pgtype.Text{String: req.GetDeviceOs(), Valid: req.DeviceOs != ""},
		DeviceRam:               pgtype.Float8{Float64: req.GetDeviceRam(), Valid: req.DeviceRam != 0},
		DeviceStorage:           pgtype.Float8{Float64: req.GetDeviceStorage(), Valid: req.DeviceStorage != 0},
	})

	if err != nil {
		return nil, conflictError("", fmt.Sprintln("Failed to update device", err), "id")
	}

	actionIds := make([]int64, 0)
	for _, v := range deviceModel.ActionIds.([]interface{}) {
		actionIds = append(actionIds, v.(int64))
	}

	rsp := &device.ActivateDeviceResponse{
		AccessToken:           token,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(payload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
		ActionIds:             actionIds,
	}

	return rsp, nil
}
