package gapi

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/nicodanke/gesty-api/services/account-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/account-service/sse"
	"github.com/nicodanke/gesty-api/services/account-service/validators"
	userValidator "github.com/nicodanke/gesty-api/services/account-service/validators/user"
	"github.com/nicodanke/gesty-api/shared/proto/account-service/requests/user"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

const (
	sse_update_user = "update-user"
)

func (server *Server) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	log.Info().Str("method", "UpdateUser").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing UpdateUser request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAU", "UU"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAU or UU"))
	}

	violations := validateUpdateUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	if req.GetRoleId() != 0 {
		getRoleParams := db.GetRoleParams{
			AccountID: authPayload.AccountID,
			ID:        req.GetRoleId(),
		}

		_, err = server.store.GetRole(ctx, getRoleParams)
		if err != nil {
			return nil, conflictError("", fmt.Sprintln("Role not found"), "role_id")
		}
	}

	arg := db.UpdateUserParams{
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
		Phone: pgtype.Text{
			String: req.GetPhone(),
			Valid:  req.Phone != nil,
		},
		Email: pgtype.Text{
			String: req.GetEmail(),
			Valid:  req.Email != nil,
		},
		RoleID: pgtype.Int8{
			Int64: req.GetRoleId(),
			Valid: req.RoleId != nil,
		},
		Active: pgtype.Bool{
			Bool:  req.GetActive(),
			Valid: req.Active != nil,
		},
		IsAdmin: pgtype.Bool{
			Bool:  req.GetIsAdmin(),
			Valid: req.IsAdmin != nil,
		},
		UpdatedAt: pgtype.Timestamptz{
			Time:  time.Now().UTC(),
			Valid: true,
		},
	}

	result, err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.UniqueViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_UNIQUE, fmt.Sprintln("Failed to update user due to unique constraint violation"), constraintName)
		}
		if errCode == db.ForeignKeyViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_FK, fmt.Sprintln("Failed to update user due to foreign key constraint violation"), constraintName)
		}

		return nil, internalError(fmt.Sprintln("Failed to update user", err))
	}

	userModel := convertUser(result)
	userEvent := convertUserEvent(result)

	rsp := &user.UpdateUserResponse{
		User: userModel,
	}

	// Notify account update
	server.notifier.BoadcastMessageToAccount(sse.NewEventMessage(sse_update_user, userEvent), result.ID, nil)

	return rsp, nil
}

func validateUpdateUserRequest(req *user.UpdateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if req.Name != nil {
		if err := userValidator.ValidateName(req.GetName()); err != nil {
			violations = append(violations, fieldViolation("name", err))
		}
	}

	if req.Lastname != nil {
		if err := userValidator.ValidateLastname(req.GetLastname()); err != nil {
			violations = append(violations, fieldViolation("lastname", err))
		}
	}

	if req.Email != nil {
		if err := validators.ValidateEmail(req.GetEmail()); err != nil {
			violations = append(violations, fieldViolation("email", err))
		}
	}

	if req.Phone != nil {
		if err := userValidator.ValidatePhone(req.GetPhone()); err != nil {
			violations = append(violations, fieldViolation("phone", err))
		}
	}

	return violations
}
