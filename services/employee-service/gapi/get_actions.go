package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/employee-service/validators"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/action"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *Server) GetActions(ctx context.Context, req *action.GetActionsRequest) (*action.GetActionsResponse, error) {
	log.Info().Str("method", "GetActions").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing GetActions request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAA", "LA"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAA or LA"))
	}

	violations := validateGetActionsRequest(req)
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

	arg := db.GetActionsParams{
		AccountID: authPayload.AccountID,
		Limit:     limit,
		Offset:    offset,
	}

	result, err := server.store.GetActions(ctx, arg)
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to get actions:", err))
	}

	rsp := &action.GetActionsResponse{
		Actions: convertActions(result),
	}
	return rsp, nil
}

func validateGetActionsRequest(req *action.GetActionsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
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
