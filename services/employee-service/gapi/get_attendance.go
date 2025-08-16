package gapi

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/attendance"
	"github.com/rs/zerolog/log"
)

func (server *Server) GetAttendance(ctx context.Context, req *attendance.GetAttendanceRequest) (*attendance.GetAttendanceResponse, error) {
	log.Info().Str("method", "GetAttendance").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing GetAttendance request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAAT", "LAT"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAAT or LAT"))
	}

	result, err := server.store.GetAttendance(ctx, db.GetAttendanceParams{
		AccountID: authPayload.AccountID,
		ID:        uuid.MustParse(req.GetId()),
	})
	if err != nil {
		if err == db.ErrRecordNotFound {
			return nil, notFoundError("", fmt.Sprintln("Attendance with id", req.GetId(), "not found"))
		}
		return nil, internalError(fmt.Sprintln("Failed to get attendance:", err))
	}

	rsp := &attendance.GetAttendanceResponse{
		Attendance: convertAttendanceRow(result),
	}
	return rsp, nil
}
