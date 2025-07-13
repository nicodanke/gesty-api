package gapi

import (
	"context"
	"fmt"
	"strconv"

	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/employee-service/sse"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/facility"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	sse_delete_facility = "delete-facility"
)

func (server *Server) DeleteFacility(ctx context.Context, req *facility.DeleteFacilityRequest) (*emptypb.Empty, error) {
	log.Info().Str("method", "DeleteFacility").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing DeleteFacility request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAF", "DF"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAF or DF"))
	}

	argDelete := db.DeleteFacilityTxParams{
		AccountID: authPayload.AccountID,
		ID:        req.GetId(),
	}

	err = server.store.DeleteFacilityTx(ctx, argDelete)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.ForeignKeyViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_FK, fmt.Sprintln("Failed to delete facility due to foreign key constraint violation"), constraintName)
		}

		return nil, internalError(fmt.Sprintln("Failed to delete facility:", err))
	}

	// Notify delete user
	var data = map[string]any{}
	data["id"] = strconv.FormatInt(req.GetId(), 10)
	server.notifier.BoadcastMessageToAccount(sse.NewEventMessage(sse_delete_facility, data), authPayload.AccountID, &authPayload.UserID)

	return &emptypb.Empty{}, nil
}
