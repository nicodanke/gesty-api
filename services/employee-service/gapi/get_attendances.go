package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/employee-service/validators"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/attendance"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *Server) GetAttendances(ctx context.Context, req *attendance.GetAttendancesRequest) (*attendance.GetAttendancesResponse, error) {
	log.Info().Str("method", "GetAttendances").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing GetAttendances request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAAT", "LAT"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAAT or LAT"))
	}

	violations := validateGetAttendancesRequest(req)
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

	arg := db.GetAttendancesParams{
		AccountID: authPayload.AccountID,
		Limit:     limit,
		Offset:    offset,
	}

	result, err := server.store.GetAttendances(ctx, arg)
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to get attendances:", err))
	}

	fmt.Println("result", result)

	rsp := &attendance.GetAttendancesResponse{
		Attendances: convertAttendancesRows(result),
	}
	return rsp, nil
}

func validateGetAttendancesRequest(req *attendance.GetAttendancesRequest) (violations []*errdetails.BadRequest_FieldViolation) {
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
