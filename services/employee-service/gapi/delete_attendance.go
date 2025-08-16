package gapi

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/employee-service/sse"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/attendance"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	sse_delete_attendance = "delete-attendance"
)

func (server *Server) DeleteAttendance(ctx context.Context, req *attendance.DeleteAttendanceRequest) (*emptypb.Empty, error) {
	log.Info().Str("method", "DeleteAttendance").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing DeleteAttendance request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAAT", "DAT"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAAT or DAT"))
	}

	arg := db.GetAttendanceParams{
		AccountID: authPayload.AccountID,
		ID:        uuid.MustParse(req.GetId()),
	}

	_, err = server.store.GetAttendance(ctx, arg)
	if err != nil {
		return nil, notFoundError(NOT_FOUND, fmt.Sprintln("Attendance not found:", err))
	}

	err = server.store.DeleteAttendance(ctx, uuid.MustParse(req.GetId()))
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.ForeignKeyViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_FK, fmt.Sprintln("Failed to delete attendance due to foreign key constraint violation"), constraintName)
		}

		return nil, internalError(fmt.Sprintln("Failed to delete attendance:", err))
	}

	// Notify delete attendace
	var data = map[string]any{}
	data["id"] = req.GetId()
	server.notifier.BoadcastMessageToAccount(sse.NewEventMessage(sse_delete_attendance, data), authPayload.AccountID, &authPayload.UserID)

	return &emptypb.Empty{}, nil
}
