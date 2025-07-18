package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/employee-service/sse"
	"github.com/nicodanke/gesty-api/services/employee-service/utils"
	deviceValidator "github.com/nicodanke/gesty-api/services/employee-service/validators/device"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/device"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

const (
	sse_create_device = "create-device"
)

func (server *Server) CreateDevice(ctx context.Context, req *device.CreateDeviceRequest) (*device.CreateDeviceResponse, error) {
	log.Info().Str("method", "CreateDevice").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing CreateDevice request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAD", "CD"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAD or CD"))
	}

	violations := validateCreateDeviceRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	_, err = server.store.GetFacility(ctx, db.GetFacilityParams{AccountID: authPayload.AccountID, ID: req.GetFacilityId()})
	if err != nil {
		return nil, conflictError("", fmt.Sprintln("Facility not found"), "facility_id")
	}

	password, err := utils.GenerateRandomString(12)
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to generate random password", err))
	}

	arg := db.CreateDeviceTxParams{
		AccountID:  authPayload.AccountID,
		Name:       req.GetName(),
		Enabled:    req.GetEnabled(),
		Password:   password,
		FacilityID: req.GetFacilityId(),
		ActionIDs:  req.GetActionIds(),
	}

	result, err := server.store.CreateDeviceTx(ctx, arg)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.UniqueViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_UNIQUE, fmt.Sprintln("Failed to create device due to unique constraint violation"), constraintName)
		}
		if errCode == db.ForeignKeyViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_FK, fmt.Sprintln("Failed to create device due to foreign key constraint violation"), constraintName)
		}
		return nil, internalError(fmt.Sprintln("Failed to create device", err))
	}

	deviceModel := convertDeviceCreateTxResult(result)
	deviceEvent := convertDeviceCreateTxResultEvent(result)

	rsp := &device.CreateDeviceResponse{
		Device: deviceModel,
	}

	// Notify device creation
	server.notifier.BoadcastMessageToAccount(sse.NewEventMessage(sse_create_device, deviceEvent), authPayload.AccountID, &authPayload.UserID)

	return rsp, nil
}

func validateCreateDeviceRequest(req *device.CreateDeviceRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := deviceValidator.ValidateName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}

	return violations
}
