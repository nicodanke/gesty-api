package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/gesty-api/services/account-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/account-service/requests/role"
	"github.com/nicodanke/gesty-api/services/account-service/sse"
	roleValidator "github.com/nicodanke/gesty-api/services/account-service/validators/role"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

const (
	sse_create_role = "create-role"
)

func (server *Server) CreateRole(ctx context.Context, req *role.CreateRoleRequest) (*role.CreateRoleResponse, error) {
	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAR", "CR"}})
	if !authorized {
		return nil, permissionDeniedError("FORBIDDEN", fmt.Sprintln("User not authorized"))
	}

	violations := validateCreateRoleRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.CreateRoleTxParams{
		AccountID:      authPayload.AccountID,
		Name:           req.GetName(),
		Description:    req.GetDescription(),
		PermissionIDs:  req.GetPermissionIds(),
	}

	result, err := server.store.CreateRoleTx(ctx, arg)
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to create role", err))
	}

	rsp := &role.CreateRoleResponse{
		Role: convertRoleCreate(result),
	}

	// Notify role creation
	server.notifier.BoadcastMessageToAccount(sse.NewEventMessage(sse_create_role, rsp), authPayload.AccountID, &authPayload.UserID)

	return rsp, nil
}

func validateCreateRoleRequest(req *role.CreateRoleRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := roleValidator.ValidateName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}

	return violations
}