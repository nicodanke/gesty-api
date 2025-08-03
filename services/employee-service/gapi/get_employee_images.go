package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/employee"
	"github.com/rs/zerolog/log"
)

func (server *Server) GetImagesEmployee(ctx context.Context, req *employee.GetImagesEmployeeRequest) (*employee.GetImagesEmployeeResponse, error) {
	log.Info().Str("method", "GetImagesEmployee").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing GetImagesEmployee request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAE", "LE"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAE or CE"))
	}

	result, err := server.store.GetEmployeePhotos(ctx, db.GetEmployeePhotosParams{
		EmployeeID: req.GetId(),
		AccountID:  authPayload.AccountID,
		Limit:      100,
		Offset:     0,
	})
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to get employee profile photo", err))
	}

	rsp := &employee.GetImagesEmployeeResponse{
		EmployeeImages: convertEmployeePhoto(result),
	}

	return rsp, nil
}
