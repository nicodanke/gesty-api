package gapi

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/nicodanke/gesty-api/services/account-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/account-service/requests/user"
	"github.com/nicodanke/gesty-api/services/account-service/sse"
	"github.com/nicodanke/gesty-api/shared/utils"
	"github.com/nicodanke/gesty-api/services/account-service/validators"
	userValidator "github.com/nicodanke/gesty-api/services/account-service/validators/user"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

const (
	sse_create_user = "create-user"
)

func (server *Server) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	violations := validateCreateUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	hashedPassword, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to hash password", err))
	}

	role, err := server.store.GetRole(ctx, db.GetRoleParams{AccountID: authPayload.AccountID, ID: req.GetRoleId()})
	if err != nil {
		violations = append(violations, fieldViolation("roleId", fmt.Errorf("error checking if role is  valid: %w", err)))
		return nil, invalidArgumentError(violations)
	}

	arg := db.CreateUserParams{
		Name:     req.GetName(),
		Lastname: req.GetLastname(),
		Username: req.GetUsername() + "@" + authPayload.AccountCode,
		Email:    req.GetEmail(),
		Password: hashedPassword,
		Phone: pgtype.Text{
			String: req.GetPhone(),
			Valid:  req.Phone != nil,
		},
		Active:    req.GetActive(),
		IsAdmin:   req.GetIsAdmin(),
		RoleID:    role.ID,
		AccountID: authPayload.AccountID,
	}

	result, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to create user", err))
	}

	rsp := &user.CreateUserResponse{
		User: convertUser(result),
	}

	// Notify user creation
	server.notifier.BoadcastMessageToAccount(sse.NewEventMessage(sse_create_user, rsp), authPayload.AccountID, &authPayload.UserID)

	return rsp, nil
}

func validateCreateUserRequest(req *user.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {

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

	if err := validators.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}

	if req.Phone != nil {
		if err := userValidator.ValidatePhone(req.GetPhone()); err != nil {
			violations = append(violations, fieldViolation("phone", err))
		}
	}

	return violations
}