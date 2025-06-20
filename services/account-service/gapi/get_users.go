package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/gesty-api/services/account-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/account-service/requests/user"
	"github.com/nicodanke/gesty-api/services/account-service/validators"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *Server) GetUsers(ctx context.Context, req *user.GetUsersRequest) (*user.GetUsersResponse, error) {
	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	violations := validateGetUsersRequest(req)
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

	arg := db.GetUsersParams{
		AccountID: authPayload.AccountID,
		Limit:     limit,
		Offset:    offset,
	}

	result, err := server.store.GetUsers(ctx, arg)
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to get users:", err))
	}

	rsp := &user.GetUsersResponse{
		Users: convertUsers(result),
	}
	return rsp, nil
}

func validateGetUsersRequest(req *user.GetUsersRequest) (violations []*errdetails.BadRequest_FieldViolation) {
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