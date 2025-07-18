package gapi

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/employee-service/sse"
	deviceValidator "github.com/nicodanke/gesty-api/services/employee-service/validators/device"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/device"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

const (
	sse_update_device = "update-device"
)

func (server *Server) UpdateDevice(ctx context.Context, req *device.UpdateDeviceRequest) (*device.UpdateDeviceResponse, error) {
	log.Info().Str("method", "UpdateDevice").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing UpdateDevice request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAD", "UD"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAD or UD"))
	}

	violations := validateUpdateDeviceRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	if req.GetFacilityId() != 0 {
		getFacilityParams := db.GetFacilityParams{
			AccountID: authPayload.AccountID,
			ID:        req.GetFacilityId(),
		}

		_, err = server.store.GetFacility(ctx, getFacilityParams)
		if err != nil {
			return nil, conflictError("", fmt.Sprintln("Facility not found"), "facility_id")
		}
	}

	actionIds := []int64{}
	if !req.GetRemoveAllActions() {
		actionIds = req.GetActionIds()
	}

	arg := db.UpdateDeviceTxParams{
		AccountID: authPayload.AccountID,
		ID:        req.GetId(),
		Name: pgtype.Text{
			String: req.GetName(),
			Valid:  req.Name != nil,
		},
		Enabled: pgtype.Bool{
			Bool:  req.GetEnabled(),
			Valid: req.Enabled != nil,
		},
		FacilityID: pgtype.Int8{
			Int64: req.GetFacilityId(),
			Valid: req.GetFacilityId() != 0,
		},
		ActionIDs: actionIds,
	}

	result, err := server.store.UpdateDeviceTx(ctx, arg)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.UniqueViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_UNIQUE, fmt.Sprintln("Failed to update device due to unique constraint violation"), constraintName)
		}
		if errCode == db.ForeignKeyViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_FK, fmt.Sprintln("Failed to update device due to foreign key constraint violation"), constraintName)
		}

		return nil, internalError(fmt.Sprintln("Failed to update device", err))
	}

	deviceModel := convertDeviceUpdateTxResult(result)
	deviceEvent := convertDeviceUpdateTxResultEvent(result)

	rsp := &device.UpdateDeviceResponse{
		Device: deviceModel,
	}

	// Notify device update
	server.notifier.BoadcastMessageToAccount(sse.NewEventMessage(sse_update_device, deviceEvent), authPayload.AccountID, &authPayload.UserID)

	return rsp, nil
}

func validateUpdateDeviceRequest(req *device.UpdateDeviceRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if req.Name != nil {
		if err := deviceValidator.ValidateName(req.GetName()); err != nil {
			violations = append(violations, fieldViolation("name", err))
		}
	}

	return violations
}
