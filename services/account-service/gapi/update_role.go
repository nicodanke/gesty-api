package gapi

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/nicodanke/gesty-api/services/account-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/account-service/requests/role"
	"github.com/nicodanke/gesty-api/services/account-service/sse"
	roleValidator "github.com/nicodanke/gesty-api/services/account-service/validators/role"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

const (
	sse_update_role = "update-role"
)

func (server *Server) UpdateRole(ctx context.Context, req *role.UpdateRoleRequest) (*role.UpdateRoleResponse, error) {
	log.Info().Str("method", "UpdateRole").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing UpdateRole request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAR", "UR"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAR or UR"))
	}

	violations := validateUpdateRoleRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.UpdateRoleTxParams{
		AccountID: authPayload.AccountID,
		ID: req.GetId(),
		Name: pgtype.Text{
			String: req.GetName(),
			Valid:  req.Name != nil,
		},
		Description: pgtype.Text{
			String: req.GetDescription(),
			Valid:  req.Description != nil,
		},
		PermissionIDs: req.GetPermissionIds(),
	}

	result, err := server.store.UpdateRoleTx(ctx, arg)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.UniqueViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_UNIQUE, fmt.Sprintln("Failed to update role due to unique constraint violation"), constraintName)
		}
		if errCode == db.ForeignKeyViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_FK, fmt.Sprintln("Failed to update role due to foreign key constraint violation"), constraintName)
		}

		return nil, internalError(fmt.Sprintln("Failed to update role", err))
	}

	rsp := &role.UpdateRoleResponse{
		Role: convertRoleUpdate(result),
	}

	// Notify account update
	server.notifier.BoadcastMessageToAccount(sse.NewEventMessage(sse_update_role, rsp), result.Role.ID, nil)

	return rsp, nil
}

func validateUpdateRoleRequest(req *role.UpdateRoleRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if req.Name != nil {
		if err := roleValidator.ValidateName(req.GetName()); err != nil {
			violations = append(violations, fieldViolation("name", err))
		}
	}

	return violations
}