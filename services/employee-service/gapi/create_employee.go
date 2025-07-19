package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/employee-service/sse"
	employeeValidator "github.com/nicodanke/gesty-api/services/employee-service/validators/employee"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/employee"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

const (
	sse_create_employee = "create-employee"
)

func (server *Server) CreateEmployee(ctx context.Context, req *employee.CreateEmployeeRequest) (*employee.CreateEmployeeResponse, error) {
	log.Info().Str("method", "CreateEmployee").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing CreateEmployee request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAE", "CE"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAE or CE"))
	}

	violations := validateCreateEmployeeRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.CreateEmployeeTxParams{
		AccountID:       authPayload.AccountID,
		Name:            req.GetName(),
		Lastname:        req.GetLastname(),
		Email:           req.GetEmail(),
		Phone:           req.GetPhone(),
		Gender:          req.GetGender(),
		RealID:          req.GetRealId(),
		FiscalID:        req.GetFiscalId(),
		AddressCountry:  req.GetAddressCountry(),
		AddressState:    req.GetAddressState(),
		AddressSubState: req.GetAddressSubState(),
		AddressStreet:   req.GetAddressStreet(),
		AddressNumber:   req.GetAddressNumber(),
		AddressUnit:     req.GetAddressUnit(),
		AddressZipCode:  req.GetAddressZipCode(),
		AddressLat:      req.GetAddressLat(),
		AddressLng:      req.GetAddressLng(),
		FacilityIDs:     req.GetFacilityIds(),
	}

	result, err := server.store.CreateEmployeeTx(ctx, arg)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.UniqueViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_UNIQUE, fmt.Sprintln("Failed to create employee due to unique constraint violation"), constraintName)
		}
		if errCode == db.ForeignKeyViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_FK, fmt.Sprintln("Failed to create employee due to foreign key constraint violation"), constraintName)
		}
		return nil, internalError(fmt.Sprintln("Failed to create action", err))
	}

	employeeModel := convertEmployeeCreateTxResult(result)
	employeeEvent := convertEmployeeCreateTxResultEvent(result)

	rsp := &employee.CreateEmployeeResponse{
		Employee: employeeModel,
	}

	// Notify role creation
	server.notifier.BoadcastMessageToAccount(sse.NewEventMessage(sse_create_employee, employeeEvent), authPayload.AccountID, &authPayload.UserID)

	return rsp, nil
}

func validateCreateEmployeeRequest(req *employee.CreateEmployeeRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := employeeValidator.ValidateName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}

	if err := employeeValidator.ValidateLastname(req.GetLastname()); err != nil {
		violations = append(violations, fieldViolation("lastname", err))
	}

	if err := employeeValidator.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}

	if err := employeeValidator.ValidatePhone(req.GetPhone()); err != nil {
		violations = append(violations, fieldViolation("phone", err))
	}

	if err := employeeValidator.ValidateGender(req.GetGender()); err != nil {
		violations = append(violations, fieldViolation("gender", err))
	}

	if err := employeeValidator.ValidateRealId(req.GetRealId()); err != nil {
		violations = append(violations, fieldViolation("realId", err))
	}

	if err := employeeValidator.ValidateFiscalId(req.GetFiscalId()); err != nil {
		violations = append(violations, fieldViolation("fiscalId", err))
	}

	if err := employeeValidator.ValidateAddressCountry(req.GetAddressCountry()); err != nil {
		violations = append(violations, fieldViolation("addressCountry", err))
	}

	if err := employeeValidator.ValidateAddressState(req.GetAddressState()); err != nil {
		violations = append(violations, fieldViolation("addressState", err))
	}

	if err := employeeValidator.ValidateAddressSubState(req.GetAddressSubState()); err != nil {
		violations = append(violations, fieldViolation("addressSubState", err))
	}

	if err := employeeValidator.ValidateAddressStreet(req.GetAddressStreet()); err != nil {
		violations = append(violations, fieldViolation("addressStreet", err))
	}

	if err := employeeValidator.ValidateAddressNumber(req.GetAddressNumber()); err != nil {
		violations = append(violations, fieldViolation("addressNumber", err))
	}

	if err := employeeValidator.ValidateAddressUnit(req.GetAddressUnit()); err != nil {
		violations = append(violations, fieldViolation("addressUnit", err))
	}

	if err := employeeValidator.ValidateAddressZipCode(req.GetAddressZipCode()); err != nil {
		violations = append(violations, fieldViolation("addressZipCode", err))
	}

	if req.GetAddressLat() != 0 {
		if err := employeeValidator.ValidateAddressLat(req.GetAddressLat()); err != nil {
			violations = append(violations, fieldViolation("addressLat", err))
		}
	}

	if req.GetAddressLng() != 0 {
		if err := employeeValidator.ValidateAddressLng(req.GetAddressLng()); err != nil {
			violations = append(violations, fieldViolation("addressLng", err))
		}
	}

	return violations
}
