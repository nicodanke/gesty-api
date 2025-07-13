package gapi

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/employee-service/sse"
	facilityValidator "github.com/nicodanke/gesty-api/services/employee-service/validators/facility"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/facility"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

const (
	sse_update_facility = "update-facility"
)

func (server *Server) UpdateFacility(ctx context.Context, req *facility.UpdateFacilityRequest) (*facility.UpdateFacilityResponse, error) {
	log.Info().Str("method", "UpdateFacility").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing UpdateFacility request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAF", "UF"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAF or UF"))
	}

	violations := validateUpdateFacilityRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.UpdateFacilityTxParams{
		AccountID: authPayload.AccountID,
		ID:        req.GetId(),
		Name: pgtype.Text{
			String: req.GetName(),
			Valid:  req.Name != nil,
		},
		Description: pgtype.Text{
			String: req.GetDescription(),
			Valid:  req.Description != nil,
		},
		OpenTime: pgtype.Time{
			Microseconds: req.GetOpenTime().AsDuration().Microseconds(),
			Valid:        req.OpenTime != nil,
		},
		CloseTime: pgtype.Time{
			Microseconds: req.GetCloseTime().AsDuration().Microseconds(),
			Valid:        req.CloseTime != nil,
		},
		AddressCountry: pgtype.Text{
			String: req.GetAddressCountry(),
			Valid:  req.AddressCountry != nil,
		},
		AddressState: pgtype.Text{
			String: req.GetAddressState(),
			Valid:  req.AddressState != nil,
		},
		AddressSubState: pgtype.Text{
			String: req.GetAddressSubState(),
			Valid:  req.AddressSubState != nil,
		},
		AddressStreet: pgtype.Text{
			String: req.GetAddressStreet(),
			Valid:  req.AddressStreet != nil,
		},
		AddressNumber: pgtype.Text{
			String: req.GetAddressNumber(),
			Valid:  req.AddressNumber != nil,
		},
		AddressUnit: pgtype.Text{
			String: req.GetAddressUnit(),
			Valid:  req.AddressUnit != nil,
		},
		AddressPostalcode: pgtype.Text{
			String: req.GetAddressPostalcode(),
			Valid:  req.AddressPostalcode != nil,
		},
		AddressLat: pgtype.Float8{
			Float64: req.GetAddressLat(),
			Valid:   req.AddressLat != nil,
		},
		AddressLng: pgtype.Float8{
			Float64: req.GetAddressLng(),
			Valid:   req.AddressLng != nil,
		},
	}

	result, err := server.store.UpdateFacilityTx(ctx, arg)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.UniqueViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_UNIQUE, fmt.Sprintln("Failed to update facility due to unique constraint violation"), constraintName)
		}
		if errCode == db.ForeignKeyViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_FK, fmt.Sprintln("Failed to update facility due to foreign key constraint violation"), constraintName)
		}

		return nil, internalError(fmt.Sprintln("Failed to update facility", err))
	}

	facilityModel := convertFacilityUpdateTxResult(result)
	facilityEvent := convertFacilityUpdateTxResult(result)

	rsp := &facility.UpdateFacilityResponse{
		Facility: facilityModel,
	}

	// Notify account update
	server.notifier.BoadcastMessageToAccount(sse.NewEventMessage(sse_update_facility, facilityEvent), authPayload.AccountID, &authPayload.UserID)

	return rsp, nil
}

func validateUpdateFacilityRequest(req *facility.UpdateFacilityRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if req.Name != nil {
		if err := facilityValidator.ValidateName(req.GetName()); err != nil {
			violations = append(violations, fieldViolation("name", err))
		}
	}

	if req.Description != nil {
		if err := facilityValidator.ValidateDescription(req.GetDescription()); err != nil {
			violations = append(violations, fieldViolation("description", err))
		}
	}

	if req.OpenTime != nil {
		if err := facilityValidator.ValidateOpenTime(req.GetOpenTime()); err != nil {
			violations = append(violations, fieldViolation("openTime", err))
		}
	}

	if req.CloseTime != nil {
		if err := facilityValidator.ValidateCloseTime(req.GetCloseTime()); err != nil {
			violations = append(violations, fieldViolation("closeTime", err))
		}
	}

	if req.AddressCountry != nil {
		if err := facilityValidator.ValidateAddressCountry(req.GetAddressCountry()); err != nil {
			violations = append(violations, fieldViolation("addressCountry", err))
		}
	}

	if req.AddressState != nil {
		if err := facilityValidator.ValidateAddressState(req.GetAddressState()); err != nil {
			violations = append(violations, fieldViolation("addressState", err))
		}
	}

	if req.AddressSubState != nil {
		if err := facilityValidator.ValidateAddressSubState(req.GetAddressSubState()); err != nil {
			violations = append(violations, fieldViolation("addressSubState", err))
		}
	}

	if req.AddressStreet != nil {
		if err := facilityValidator.ValidateAddressStreet(req.GetAddressStreet()); err != nil {
			violations = append(violations, fieldViolation("addressStreet", err))
		}
	}

	if req.AddressNumber != nil {
		if err := facilityValidator.ValidateAddressNumber(req.GetAddressNumber()); err != nil {
			violations = append(violations, fieldViolation("addressNumber", err))
		}
	}

	if req.AddressUnit != nil {
		if err := facilityValidator.ValidateAddressUnit(req.GetAddressUnit()); err != nil {
			violations = append(violations, fieldViolation("addressUnit", err))
		}
	}

	if req.AddressPostalcode != nil {
		if err := facilityValidator.ValidateAddressPostalcode(req.GetAddressPostalcode()); err != nil {
			violations = append(violations, fieldViolation("addressPostalcode", err))
		}
	}

	if req.AddressLat != nil {
		if err := facilityValidator.ValidateAddressLat(*req.AddressLat); err != nil {
			violations = append(violations, fieldViolation("addressLat", err))
		}
	}

	if req.AddressLng != nil {
		if err := facilityValidator.ValidateAddressLng(*req.AddressLng); err != nil {
			violations = append(violations, fieldViolation("addressLng", err))
		}
	}

	return violations
}
