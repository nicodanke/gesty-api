package gapi

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/employee-service/utils"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/device"
	"github.com/rs/zerolog/log"
)

func (server *Server) GenerateActivationCode(ctx context.Context, req *device.GenerateActivationCodeRequest) (*device.GenerateActivationCodeResponse, error) {
	log.Info().Str("method", "GenerateDeviceActivationCode").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing GenerateDeviceActivationCode request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAD", "CACD"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAD or CACD"))
	}

	_, err = server.store.GetDevice(ctx, db.GetDeviceParams{AccountID: authPayload.AccountID, ID: req.GetId()})
	if err != nil {
		return nil, conflictError("", fmt.Sprintln("Device not found"), "id")
	}

	deviceModel, err := server.store.GetDevice(ctx, db.GetDeviceParams{AccountID: authPayload.AccountID, ID: req.GetId()})
	if err != nil {
		return nil, conflictError("", fmt.Sprintln("Device not found"), "id")
	}

	if !deviceModel.Enabled {
		return nil, conflictError("", fmt.Sprintln("Device is not enabled"), "id")
	}

	activationCode, err := utils.GenerateRandomString(36)
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to generate random activation code", err))
	}

	_, err = server.store.UpdateDevice(ctx, db.UpdateDeviceParams{
		AccountID: authPayload.AccountID,
		ID:        req.GetId(),
		ActivationCode: pgtype.Text{
			String: activationCode,
			Valid:  true,
		},
		ActivationCodeExpiresAt: pgtype.Timestamptz{
			Time:  time.Now().Add(5 * time.Minute),
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		},
	})

	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.UniqueViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_UNIQUE, fmt.Sprintln("Failed to generate activation code due to unique constraint violation"), constraintName)
		}
		if errCode == db.ForeignKeyViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_FK, fmt.Sprintln("Failed to generate activation code due to foreign key constraint violation"), constraintName)
		}
		return nil, internalError(fmt.Sprintln("Failed to generate activation code", err))
	}

	code := strconv.FormatInt(authPayload.AccountID, 10) + "|" + strconv.FormatInt(deviceModel.ID, 10) + "|" + activationCode

	base64Code := base64.StdEncoding.EncodeToString([]byte(code))

	rsp := &device.GenerateActivationCodeResponse{
		Code: base64Code,
	}
	return rsp, nil
}
