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
	sse_update_action = "update-action"
)

func (server *Server) UpdateAction(ctx context.Context, req *action.UpdateActionRequest) (*action.UpdateActionResponse, error) {
	log.Info().Str("method", "UpdateAction").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing UpdateAction request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"SAA", "UA"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: SAA or UA"))
	}

	violations := validateUpdateActionRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.UpdateActionParams{
		AccountID: authPayload.AccountID,
		ID:        req.GetId(),
		Name: pgtype.Text{
			String: req.GetName(),
			Valid:  req.Name != nil,
		},
		Description: pgtype.Text{
			String: req.GetDescription(),
			Valid:  req.Description != nil,
		},
		Enabled: pgtype.Bool{
			Bool:  req.GetEnabled(),
			Valid: req.Enabled != nil,
		},
	}

	result, err := server.store.UpdateAction(ctx, arg)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.UniqueViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_UNIQUE, fmt.Sprintln("Failed to update action due to unique constraint violation"), constraintName)
		}
		if errCode == db.ForeignKeyViolation {
			constraintName := db.ConstraintName(err)
			return nil, conflictError(CONFLICT_FK, fmt.Sprintln("Failed to update action due to foreign key constraint violation"), constraintName)
		}

		return nil, internalError(fmt.Sprintln("Failed to update action", err))
	}

	actionModel := convertAction(result)
	actionEvent := convertActionEvent(result)

	rsp := &action.UpdateActionResponse{
		Action: actionModel,
	}

	// Notify account update
	server.notifier.BoadcastMessageToAccount(sse.NewEventMessage(sse_update_action, actionEvent), authPayload.AccountID, &authPayload.UserID)

	return rsp, nil
}

func validateUpdateActionRequest(req *action.UpdateActionRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if req.Name != nil {
		if err := actionValidator.ValidateName(req.GetName()); err != nil {
			violations = append(violations, fieldViolation("name", err))
		}
	}

	if req.Description != nil {
		if err := actionValidator.ValidateDescription(req.GetDescription()); err != nil {
			violations = append(violations, fieldViolation("description", err))
		}
	}

	return violations
}
