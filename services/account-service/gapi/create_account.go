package gapi

import (
	"context"
	"fmt"
	"strings"

	db "github.com/nicodanke/gesty-api/services/account-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/account-service/requests/account"
	"github.com/nicodanke/gesty-api/shared/utils"
	"github.com/nicodanke/gesty-api/services/account-service/validators"
	accountValidator "github.com/nicodanke/gesty-api/services/account-service/validators/account"
	userValidator "github.com/nicodanke/gesty-api/services/account-service/validators/user"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *Server) CreateAccount(ctx context.Context, req *account.CreateAccountRequest) (*account.CreateAccountResponse, error) {
	log.Info().Str("method", "CreateAccount").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing CreateAccount request")

	violations := validateCreateAccountRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	hashedPassword, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to hash password:", err))
	}

	code := strings.ReplaceAll(req.GetCompanyName(), " ", "")
	arg := db.CreateAccountTxParams{
		Code:           code,
		CompanyName:    req.GetCompanyName(),
		Email:          req.GetEmail(),
		Name:           req.GetName(),
		Lastname:       req.GetLastname(),
		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
	}

	result, err := server.store.CreateAccountTx(ctx, arg)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.UniqueViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_UNIQUE,fmt.Sprintln("Failed to create account due to unique constraint violation"), constraintName)
		}
		if errCode == db.ForeignKeyViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_FK, fmt.Sprintln("Failed to create account due to foreign key constraint violation"), constraintName)
		}

		return nil, internalError(fmt.Sprintln("Failed to create account:", err))
	}

	rsp := &account.CreateAccountResponse{
		Account: convertAccount(result.Account),
		User:    convertUser(result.User),
	}
	return rsp, nil
}

func validateCreateAccountRequest(req *account.CreateAccountRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := userValidator.ValidateName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}

	if err := userValidator.ValidateLastname(req.GetLastname()); err != nil {
		violations = append(violations, fieldViolation("lastname", err))
	}

	if err := userValidator.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if err := userValidator.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	if err := validators.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}

	if err := accountValidator.ValidateCompanyName(req.GetCompanyName()); err != nil {
		violations = append(violations, fieldViolation("companyName", err))
	}

	return violations
}
