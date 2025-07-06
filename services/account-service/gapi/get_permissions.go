package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/gesty-api/services/account-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/account-service/validators"
	"github.com/nicodanke/gesty-api/shared/proto/account-service/requests/permission"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *Server) GetPermissions(ctx context.Context, req *permission.GetPermissionsRequest) (*permission.GetPermissionsResponse, error) {
	log.Info().Str("method", "GetPermissions").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing GetPermissions request")

	_, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	violations := validateGetPermissionsRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	limit := int32(50)
	if req.Size != nil {
		limit = req.GetSize()
	}

	offset := int32(0)
	if req.Page != nil {
		offset = req.GetPage() * limit
	}

	arg := db.GetPermissionsParams{
		Limit:  limit,
		Offset: offset,
	}

	result, err := server.store.GetPermissions(ctx, arg)
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to get permissions:", err))
	}

	rsp := &permission.GetPermissionsResponse{
		Permissions: convertPermissions(result),
	}
	return rsp, nil
}

func validateGetPermissionsRequest(req *permission.GetPermissionsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if req.Page != nil {
		if err := validators.ValidatePage(req.GetPage()); err != nil {
			violations = append(violations, fieldViolation("page", err))
		}
	}

	if req.Size != nil {
		if err := validators.ValidateSize(req.GetSize()); err != nil {
			violations = append(violations, fieldViolation("size", err))
		}
	}

	return violations
}
