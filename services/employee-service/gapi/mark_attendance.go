package gapi

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/attendance"
	"github.com/rs/zerolog/log"
)

type MatchResponse struct {
	EmployeeID string  `json:"employee_id"`
	Precision  float64 `json:"precision"`
}

func (server *Server) MarkAttendance(ctx context.Context, req *attendance.MarkAttendanceRequest) (*attendance.MarkAttendanceResponse, error) {
	log.Info().Str("method", "MarkAttendance").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing MarkAttendance request")

	authPayload, err := server.authenticateDevice(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	device, err := server.store.GetDevice(ctx, db.GetDeviceParams{
		AccountID: authPayload.AccountID,
		ID:        authPayload.DeviceID,
	})
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to get device", err))
	}

	if !device.Enabled {
		return nil, conflictError("", fmt.Sprintln("Device is not enabled"), "id")
	}

	action, err := server.store.GetAction(ctx, db.GetActionParams{
		AccountID: authPayload.AccountID,
		ID:        req.GetActionId(),
	})
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to get action", err))
	}

	matchResponse, err := server.deepfaceClient.PostJSON("/verify", map[string]string{
		"img":        req.GetImageBase64(),
		"account_id": strconv.FormatInt(authPayload.AccountID, 10),
	})
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to register attendance", err))
	}
	defer matchResponse.Body.Close()

	// 3. Decode the body
	var result MatchResponse

	if matchResponse.StatusCode != 200 {
		fmt.Println("AAAAA", matchResponse.Status)
		return nil, conflictError("", fmt.Sprintln("No employee matched", matchResponse.Status), "id")
	}

	err = json.NewDecoder(matchResponse.Body).Decode(&result)
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to decode embedding response", err))
	}

	if result.Precision < 0.95 {
		return nil, conflictError("", fmt.Sprintln("No employee matched"), "id")
	}

	employeeId, err := strconv.ParseInt(result.EmployeeID, 10, 64)

	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to convert employee id to int64", err))
	}

	arg := db.CreateAttendanceParams{
		ID:         uuid.New(),
		TimeIn:     time.Now(),
		AccountID:  authPayload.AccountID,
		EmployeeID: employeeId,
		ActionID:   action.ID,
		DeviceID:   pgtype.Int8{Int64: device.ID, Valid: true},
		Precision:  pgtype.Float8{Float64: result.Precision, Valid: true},
	}

	attendanceResult, err := server.store.CreateAttendance(ctx, arg)
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

	employee, err := server.store.GetEmployee(ctx, db.GetEmployeeParams{
		AccountID: authPayload.AccountID,
		ID:        attendanceResult.EmployeeID,
	})
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to get employee", err))
	}

	rsp := &attendance.MarkAttendanceResponse{
		EmployeeName: fmt.Sprintf("%s %s", employee.Name, employee.Lastname),
	}

	return rsp, nil
}
