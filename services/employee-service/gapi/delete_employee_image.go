package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/employee"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) DeleteImageEmployee(ctx context.Context, req *employee.DeleteImageEmployeeRequest) (*emptypb.Empty, error) {
	log.Info().Str("method", "DeleteImageEmployee").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing DeleteImageEmployee request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAE", "LE"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAE or CE"))
	}

	err = server.store.DeleteEmployeePhoto(ctx, db.DeleteEmployeePhotoParams{
		ID:         req.GetId(),
		EmployeeID: req.GetEmployeeId(),
	})
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to delete employee photo", err))
	}

	rsp := &emptypb.Empty{}

	return rsp, nil
}
