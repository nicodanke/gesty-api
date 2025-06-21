package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/gesty-api/services/account-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/account-service/requests/role"
	"github.com/rs/zerolog/log"
)

func (server *Server) GetRole(ctx context.Context, req *role.GetRoleRequest) (*role.GetRoleResponse, error) {
	log.Info().Str("method", "GetRole").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing GetRole request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAR", "LR"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAR or LR"))
	}

	arg := db.GetRoleParams{
		AccountID: authPayload.AccountID,
		ID:        req.GetId(),
	}

	result, err := server.store.GetRole(ctx, arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return nil, notFoundError("", fmt.Sprintln("Role with id", req.GetId(), "not found"))
		}
		return nil, internalError(fmt.Sprintln("Failed to get role:", err))
	}

	rsp := &role.GetRoleResponse{
		Role: convertRoleRow(result),
	}
	return rsp, nil
}
