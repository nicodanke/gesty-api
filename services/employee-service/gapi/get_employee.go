package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/employee"
	"github.com/rs/zerolog/log"
)

func (server *Server) GetEmployee(ctx context.Context, req *employee.GetEmployeeRequest) (*employee.GetEmployeeResponse, error) {
	log.Info().Str("method", "GetEmployee").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing GetEmployee request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	if req.GetId() != authPayload.UserID {
		authorized := server.authorizeUser(authPayload, [][]string{{"SAE", "LE"}})
		if !authorized {
			return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAE or LE"))
		}
	}

	arg := db.GetEmployeeParams{
		AccountID: authPayload.AccountID,
		ID:        req.GetId(),
	}

	result, err := server.store.GetEmployee(ctx, arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return nil, notFoundError("", fmt.Sprintln("Employee with id", req.GetId(), "not found"))
		}
		return nil, internalError(fmt.Sprintln("Failed to get employee:", err))
	}

	rsp := &employee.GetEmployeeResponse{
		Employee: convertEmployeeGetRow(result),
	}
	return rsp, nil
}
