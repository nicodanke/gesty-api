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
	"github.com/rs/zerolog/log"
)

const (
	sse_create_user = "create-user"
)

func (server *Server) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	log.Info().Str("method", "CreateUser").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing CreateUser request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAU", "CU"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAU or CU"))
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
		return nil, conflictError("", fmt.Sprintln("Role not found"), "role_id")
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
		errCode := db.ErrorCode(err)
		if errCode == db.UniqueViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_UNIQUE, fmt.Sprintln("Failed to create user due to unique constraint violation"), constraintName)
		}
		if errCode == db.ForeignKeyViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_FK, fmt.Sprintln("Failed to create user due to foreign key constraint violation"), constraintName)
		}

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