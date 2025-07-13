package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/employee-service/validators"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/facility"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *Server) GetFacilities(ctx context.Context, req *facility.GetFacilitiesRequest) (*facility.GetFacilitiesResponse, error) {
	log.Info().Str("method", "GetFacilities").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing GetFacilities request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAF", "LF"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAF or LF"))
	}

	violations := validateGetFacilitiesRequest(req)
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

	arg := db.GetFacilitiesParams{
		AccountID: authPayload.AccountID,
		Limit:     limit,
		Offset:    offset,
	}

	result, err := server.store.GetFacilities(ctx, arg)
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to get facilities:", err))
	}

	rsp := &facility.GetFacilitiesResponse{
		Facilities: convertFacilitiesGetRows(result),
	}
	return rsp, nil
}

func validateGetFacilitiesRequest(req *facility.GetFacilitiesRequest) (violations []*errdetails.BadRequest_FieldViolation) {
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
