package gapi

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/employee-service/sse"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/attendance"
	"github.com/rs/zerolog/log"
)

const (
	sse_create_attendance = "create-attendance"
)

func (server *Server) CreateAttendance(ctx context.Context, req *attendance.CreateAttendanceRequest) (*attendance.CreateAttendanceResponse, error) {
	log.Info().Str("method", "CreateAttendance").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing CreateAttendance request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAAT", "CAT"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAAT or CAT"))
	}

	employee, err := server.store.GetEmployee(ctx, db.GetEmployeeParams{
		AccountID: authPayload.AccountID,
		ID:        req.GetEmployeeId(),
	})
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to get employee", err))
	}

	action, err := server.store.GetAction(ctx, db.GetActionParams{
		AccountID: authPayload.AccountID,
		ID:        req.GetActionId(),
	})
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to get action", err))
	}

	arg := db.CreateAttendanceParams{
		ID:         uuid.New(),
		TimeIn:     req.GetTimeIn().AsTime(),
		AccountID:  authPayload.AccountID,
		EmployeeID: employee.ID,
		ActionID:   action.ID,
		DeviceID:   pgtype.Int8{Int64: 1, Valid: false},
	}

	result, err := server.store.CreateAttendance(ctx, arg)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.UniqueViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_UNIQUE, fmt.Sprintln("Failed to create attendance due to unique constraint violation"), constraintName)
		}
		if errCode == db.ForeignKeyViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_FK, fmt.Sprintln("Failed to create attendance due to foreign key constraint violation"), constraintName)
		}
		return nil, internalError(fmt.Sprintln("Failed to create attendance", err))
	}

	attendanceModel := convertAttendance(result, employee, action)
	attendanceEvent := convertAttendanceEvent(result, employee, action)

	rsp := &attendance.CreateAttendanceResponse{
		Attendance: attendanceModel,
	}

	// Notify attendance creation
	server.notifier.BoadcastMessageToAccount(sse.NewEventMessage(sse_create_attendance, attendanceEvent), authPayload.AccountID, &authPayload.UserID)

	return rsp, nil
}
