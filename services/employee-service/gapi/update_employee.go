package gapi

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/employee-service/sse"
	employeeValidator "github.com/nicodanke/gesty-api/services/employee-service/validators/employee"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/employee"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

const (
	sse_update_employee = "update-employee"
)

func (server *Server) UpdateEmployee(ctx context.Context, req *employee.UpdateEmployeeRequest) (*employee.UpdateEmployeeResponse, error) {
	log.Info().Str("method", "UpdateEmployee").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing UpdateEmployee request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAE", "UE"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAE or UE"))
	}

	violations := validateUpdateEmployeeRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	facilityIds := []int64{}
	if !req.GetRemoveAllFacilities() {
		facilityIds = req.GetFacilityIds()
	}

	arg := db.UpdateEmployeeTxParams{
		AccountID: authPayload.AccountID,
		ID:        req.GetId(),
		Name: pgtype.Text{
			String: req.GetName(),
			Valid:  req.Name != nil,
		},
		Lastname: pgtype.Text{
			String: req.GetLastname(),
			Valid:  req.Lastname != nil,
		},
		Email: pgtype.Text{
			String: req.GetEmail(),
			Valid:  req.Email != nil,
		},
		Phone: pgtype.Text{
			String: req.GetPhone(),
			Valid:  req.Phone != nil,
		},
		Gender: pgtype.Text{
			String: req.GetGender(),
			Valid:  req.Gender != nil,
		},
		RealId: pgtype.Text{
			String: req.GetRealId(),
			Valid:  req.RealId != nil,
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
		AddressZipCode: pgtype.Text{
			String: req.GetAddressZipCode(),
			Valid:  req.AddressZipCode != nil,
		},
		AddressLat: pgtype.Float8{
			Float64: req.GetAddressLat(),
			Valid:   req.AddressLat != nil,
		},
		AddressLng: pgtype.Float8{
			Float64: req.GetAddressLng(),
			Valid:   req.AddressLng != nil,
		},
		FacilityIds: facilityIds,
	}

	result, err := server.store.UpdateEmployeeTx(ctx, arg)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.UniqueViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_UNIQUE, fmt.Sprintln("Failed to update employee due to unique constraint violation"), constraintName)
		}
		if errCode == db.ForeignKeyViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_FK, fmt.Sprintln("Failed to update employee due to foreign key constraint violation"), constraintName)
		}

		return nil, internalError(fmt.Sprintln("Failed to update employee", err))
	}

	employeeModel := convertEmployeeUpdateTxResult(result)
	employeeEvent := convertEmployeeUpdateTxResultEvent(result)

	rsp := &employee.UpdateEmployeeResponse{
		Employee: employeeModel,
	}

	// Notify account update
	server.notifier.BoadcastMessageToAccount(sse.NewEventMessage(sse_update_employee, employeeEvent), authPayload.AccountID, &authPayload.UserID)

	return rsp, nil
}

func validateUpdateEmployeeRequest(req *employee.UpdateEmployeeRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if req.Name != nil {
		if err := employeeValidator.ValidateName(req.GetName()); err != nil {
			violations = append(violations, fieldViolation("name", err))
		}
	}

	if req.Lastname != nil {
		if err := employeeValidator.ValidateLastname(req.GetLastname()); err != nil {
			violations = append(violations, fieldViolation("lastname", err))
		}
	}

	if req.Email != nil {
		if err := employeeValidator.ValidateEmail(req.GetEmail()); err != nil {
			violations = append(violations, fieldViolation("email", err))
		}
	}

	if req.Phone != nil {
		if err := employeeValidator.ValidatePhone(req.GetPhone()); err != nil {
			violations = append(violations, fieldViolation("phone", err))
		}
	}

	if req.Gender != nil {
		if err := employeeValidator.ValidateGender(req.GetGender()); err != nil {
			violations = append(violations, fieldViolation("gender", err))
		}
	}

	if req.RealId != nil {
		if err := employeeValidator.ValidateRealId(req.GetRealId()); err != nil {
			violations = append(violations, fieldViolation("realId", err))
		}
	}

	if req.FiscalId != nil {
		if err := employeeValidator.ValidateFiscalId(req.GetFiscalId()); err != nil {
			violations = append(violations, fieldViolation("fiscalId", err))
		}
	}

	if req.AddressCountry != nil {
		if err := employeeValidator.ValidateAddressCountry(req.GetAddressCountry()); err != nil {
			violations = append(violations, fieldViolation("addressCountry", err))
		}
	}

	if req.AddressState != nil {
		if err := employeeValidator.ValidateAddressState(req.GetAddressState()); err != nil {
			violations = append(violations, fieldViolation("addressState", err))
		}
	}

	if req.AddressSubState != nil {
		if err := employeeValidator.ValidateAddressSubState(req.GetAddressSubState()); err != nil {
			violations = append(violations, fieldViolation("addressSubState", err))
		}
	}

	if req.AddressStreet != nil {
		if err := employeeValidator.ValidateAddressStreet(req.GetAddressStreet()); err != nil {
			violations = append(violations, fieldViolation("addressStreet", err))
		}
	}

	if req.AddressNumber != nil {
		if err := employeeValidator.ValidateAddressNumber(req.GetAddressNumber()); err != nil {
			violations = append(violations, fieldViolation("addressNumber", err))
		}
	}

	if req.AddressUnit != nil {
		if err := employeeValidator.ValidateAddressUnit(req.GetAddressUnit()); err != nil {
			violations = append(violations, fieldViolation("addressUnit", err))
		}
	}

	if req.AddressZipCode != nil {
		if err := employeeValidator.ValidateAddressZipCode(req.GetAddressZipCode()); err != nil {
			violations = append(violations, fieldViolation("addressZipCode", err))
		}
	}

	if req.AddressLat != nil {
		if err := employeeValidator.ValidateAddressLat(*req.AddressLat); err != nil {
			violations = append(violations, fieldViolation("addressLat", err))
		}
	}

	if req.AddressLng != nil {
		if err := employeeValidator.ValidateAddressLng(*req.AddressLng); err != nil {
			violations = append(violations, fieldViolation("addressLng", err))
		}
	}

	return violations
}
