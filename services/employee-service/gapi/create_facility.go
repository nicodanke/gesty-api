package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/employee-service/sse"
	facilityValidator "github.com/nicodanke/gesty-api/services/employee-service/validators/facility"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/facility"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

const (
	sse_create_facility = "create-facility"
)

func (server *Server) CreateFacility(ctx context.Context, req *facility.CreateFacilityRequest) (*facility.CreateFacilityResponse, error) {
	log.Info().Str("method", "CreateFacility").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing CreateFacility request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAF", "CF"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAR or CR"))
	}

	violations := validateCreateFacilityRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.CreateFacilityTxParams{
		AccountID:         authPayload.AccountID,
		Name:              req.GetName(),
		Description:       req.GetDescription(),
		OpenTime:          req.GetOpenTime(),
		CloseTime:         req.GetCloseTime(),
		AddressCountry:    req.GetAddressCountry(),
		AddressState:      req.GetAddressState(),
		AddressSubState:   req.GetAddressSubState(),
		AddressStreet:     req.GetAddressStreet(),
		AddressNumber:     req.GetAddressNumber(),
		AddressUnit:       req.GetAddressUnit(),
		AddressPostalcode: req.GetAddressPostalcode(),
		AddressLat:        req.GetAddressLat(),
		AddressLng:        req.GetAddressLng(),
	}

	result, err := server.store.CreateFacilityTx(ctx, arg)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.UniqueViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_UNIQUE, fmt.Sprintln("Failed to create facility due to unique constraint violation"), constraintName)
		}
		if errCode == db.ForeignKeyViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_FK, fmt.Sprintln("Failed to create facility due to foreign key constraint violation"), constraintName)
		}
		return nil, internalError(fmt.Sprintln("Failed to create action", err))
	}

	facilityModel := convertFacilityCreateTxResult(result)
	facilityEvent := convertCreateFacilityTxResultEvent(result)

	rsp := &facility.CreateFacilityResponse{
		Facility: facilityModel,
	}

	// Notify role creation
	server.notifier.BoadcastMessageToAccount(sse.NewEventMessage(sse_create_facility, facilityEvent), authPayload.AccountID, &authPayload.UserID)

	return rsp, nil
}

func validateCreateFacilityRequest(req *facility.CreateFacilityRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := facilityValidator.ValidateName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}

	if err := facilityValidator.ValidateDescription(req.GetDescription()); err != nil {
		violations = append(violations, fieldViolation("description", err))
	}

	if req.GetOpenTime() != nil {
		if err := facilityValidator.ValidateOpenTime(req.GetOpenTime()); err != nil {
			violations = append(violations, fieldViolation("openTime", err))
		}
	}

	if req.GetCloseTime() != nil {
		if err := facilityValidator.ValidateCloseTime(req.GetCloseTime()); err != nil {
			violations = append(violations, fieldViolation("closeTime", err))
		}
	}

	if err := facilityValidator.ValidateAddressCountry(req.GetAddressCountry()); err != nil {
		violations = append(violations, fieldViolation("addressCountry", err))
	}

	if err := facilityValidator.ValidateAddressState(req.GetAddressState()); err != nil {
		violations = append(violations, fieldViolation("addressState", err))
	}

	if err := facilityValidator.ValidateAddressSubState(req.GetAddressSubState()); err != nil {
		violations = append(violations, fieldViolation("addressSubState", err))
	}

	if err := facilityValidator.ValidateAddressStreet(req.GetAddressStreet()); err != nil {
		violations = append(violations, fieldViolation("addressStreet", err))
	}

	if err := facilityValidator.ValidateAddressNumber(req.GetAddressNumber()); err != nil {
		violations = append(violations, fieldViolation("addressNumber", err))
	}

	if err := facilityValidator.ValidateAddressUnit(req.GetAddressUnit()); err != nil {
		violations = append(violations, fieldViolation("addressUnit", err))
	}

	if err := facilityValidator.ValidateAddressPostalcode(req.GetAddressPostalcode()); err != nil {
		violations = append(violations, fieldViolation("addressPostalcode", err))
	}

	if req.GetAddressLat() != 0 {
		if err := facilityValidator.ValidateAddressLat(req.GetAddressLat()); err != nil {
			violations = append(violations, fieldViolation("addressLat", err))
		}
	}

	if req.GetAddressLng() != 0 {
		if err := facilityValidator.ValidateAddressLng(req.GetAddressLng()); err != nil {
			violations = append(violations, fieldViolation("addressLng", err))
		}
	}

	return violations
}
