package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/employee-service/validators"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/device"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *Server) GetDevices(ctx context.Context, req *device.GetDevicesRequest) (*device.GetDevicesResponse, error) {
	log.Info().Str("method", "GetDevices").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing GetDevices request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAD", "LD"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAD or LD"))
	}

	violations := validateGetDevicesRequest(req)
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

	arg := db.GetDevicesParams{
		AccountID: authPayload.AccountID,
		Limit:     limit,
		Offset:    offset,
	}

	result, err := server.store.GetDevices(ctx, arg)
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to get devices:", err))
	}

	rsp := &device.GetDevicesResponse{
		Devices: convertGetDevicesRows(result),
	}
	return rsp, nil
}

func validateGetDevicesRequest(req *device.GetDevicesRequest) (violations []*errdetails.BadRequest_FieldViolation) {
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
