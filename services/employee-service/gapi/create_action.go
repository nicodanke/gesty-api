package gapi

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/employee-service/sse"
	actionValidator "github.com/nicodanke/gesty-api/services/employee-service/validators/action"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/action"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

const (
	sse_create_action = "create-action"
)

func (server *Server) CreateAction(ctx context.Context, req *action.CreateActionRequest) (*action.CreateActionResponse, error) {
	log.Info().Str("method", "CreateAction").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing CreateAction request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAA", "CA"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAR or CR"))
	}

	violations := validateCreateActionRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.CreateActionParams{
		AccountID:    authPayload.AccountID,
		Name:         req.GetName(),
		Description:  pgtype.Text{String: req.GetDescription(), Valid: true},
		Enabled:      req.GetEnabled(),
		CanBeDeleted: true,
	}

	result, err := server.store.CreateAction(ctx, arg)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.UniqueViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_UNIQUE, fmt.Sprintln("Failed to create action due to unique constraint violation"), constraintName)
		}
		if errCode == db.ForeignKeyViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_FK, fmt.Sprintln("Failed to create action due to foreign key constraint violation"), constraintName)
		}
		return nil, internalError(fmt.Sprintln("Failed to create action", err))
	}

	actionModel := convertAction(result)
	actionEvent := convertActionEvent(result)

	rsp := &action.CreateActionResponse{
		Action: actionModel,
	}

	// Notify role creation
	server.notifier.BoadcastMessageToAccount(sse.NewEventMessage(sse_create_action, actionEvent), authPayload.AccountID, &authPayload.UserID)

	return rsp, nil
}

func validateCreateActionRequest(req *action.CreateActionRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := actionValidator.ValidateName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}

	return violations
}
