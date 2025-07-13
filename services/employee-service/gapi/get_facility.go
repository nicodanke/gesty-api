package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/facility"
	"github.com/rs/zerolog/log"
)

func (server *Server) GetFacility(ctx context.Context, req *facility.GetFacilityRequest) (*facility.GetFacilityResponse, error) {
	log.Info().Str("method", "GetFacility").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing GetFacility request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	if req.GetId() != authPayload.UserID {
		authorized := server.authorizeUser(authPayload, [][]string{{"SAF", "LF"}})
		if !authorized {
			return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAR or CR"))
		}
	}

	arg := db.GetFacilityParams{
		AccountID: authPayload.AccountID,
		ID:        req.GetId(),
	}

	result, err := server.store.GetFacility(ctx, arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return nil, notFoundError("", fmt.Sprintln("Facility with id", req.GetId(), "not found"))
		}
		return nil, internalError(fmt.Sprintln("Failed to get facility:", err))
	}

	rsp := &facility.GetFacilityResponse{
		Facility: convertFacilityGetRow(result),
	}
	return rsp, nil
}
