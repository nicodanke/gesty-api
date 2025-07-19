package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/requests/account"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) CreateAccount(ctx context.Context, req *account.CreateAccountRequest) (*emptypb.Empty, error) {
	log.Info().Str("method", "CreateAccount").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing CreateAccount request")

	err := server.store.CreateAccountTx(ctx, db.CreateAccountTxParams{
		AccountID: req.GetAccountId(),
	})

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

	return &emptypb.Empty{}, nil
}
