package gapi

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/nicodanke/gesty-api/services/account-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/account-service/requests/account"
	"github.com/nicodanke/gesty-api/services/account-service/sse"
	"github.com/nicodanke/gesty-api/services/account-service/validators"
	accountValidator "github.com/nicodanke/gesty-api/services/account-service/validators/account"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	sse_update_account = "update-account"
)

func (server *Server) UpdateAccount(ctx context.Context, req *account.UpdateAccountRequest) (*account.UpdateAccountResponse, error) {
	violations := validateUpdateAccountRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.UpdateAccountParams{
		ID: req.GetId(),
		CompanyName: pgtype.Text{
			String: req.GetCompanyName(),
			Valid:  req.CompanyName != nil,
		},
		Email: pgtype.Text{
			String: req.GetEmail(),
			Valid:  req.Email != nil,
		},
		Phone: pgtype.Text{
			String: req.GetPhone(),
			Valid:  req.Phone != nil,
		},
		WebUrl: pgtype.Text{
			String: req.GetWebUrl(),
			Valid:  req.WebUrl != nil,
		},
		Active: pgtype.Bool{
			Bool:  req.GetActive(),
			Valid: req.Active != nil,
		},
		UpdatedAt: pgtype.Timestamptz{
			Time: time.Now().UTC(),
		},
	}

	result, err := server.store.UpdateAccount(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Fail to update account: %s", err)
	}

	rsp := &account.UpdateAccountResponse{
		Account: convertAccount(result),
	}

	// Notify account update
	server.notifier.BoadcastMessageToAccount(sse.NewEventMessage(sse_update_account, rsp), result.ID, nil)

	return rsp, nil
}

func validateUpdateAccountRequest(req *account.UpdateAccountRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if req.CompanyName != nil {
		if err := accountValidator.ValidateCompanyName(req.GetCompanyName()); err != nil {
			violations = append(violations, fieldViolation("companyName", err))
		}
	}

	if req.Email != nil {
		if err := validators.ValidateEmail(req.GetEmail()); err != nil {
			violations = append(violations, fieldViolation("email", err))
		}
	}

	if req.WebUrl != nil {
		if err := accountValidator.ValidateWebUrl(req.GetWebUrl()); err != nil {
			violations = append(violations, fieldViolation("webUrl", err))
		}
	}

	if req.Phone != nil {
		if err := accountValidator.ValidatePhone(req.GetPhone()); err != nil {
			violations = append(violations, fieldViolation("phone", err))
		}
	}

	return violations
}