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
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateAccount(ctx context.Context, req *account.CreateAccountRequest) (*account.CreateAccountResponse, error) {
	fmt.Println("CreateAccount")
	violations := validateCreateAccountRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	hashedPassword, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to hash password: %s", err)
	}

	code := strings.ReplaceAll(req.GetCompanyName(), " ", "")

	arg := db.CreateAccountTxParams{
		Code:           code,
		CompanyName:    req.GetCompanyName(),
		Email:          req.GetEmail(),
		Name:           req.GetName(),
		Lastname:       req.GetLastname(),
		Username:       req.GetUsername(),
		Country:        req.GetCountry(),
		HashedPassword: hashedPassword,
	}

	result, err := server.store.CreateAccountTx(ctx, arg)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.UniqueViolation {
			return nil, status.Error(codes.Internal, "Failed to create account: code already in use")
		}
		return nil, status.Errorf(codes.Internal, "Fail to create account: %s", err)
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

	if err := accountValidator.ValidateCountry(req.GetCountry()); err != nil {
		violations = append(violations, fieldViolation("country", err))
	}

	return violations
}
