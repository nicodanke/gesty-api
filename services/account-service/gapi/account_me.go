package gapi

import (
	"context"
	"fmt"

	"github.com/nicodanke/gesty-api/shared/proto/account-service/requests/account"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) AccountMe(ctx context.Context, req *emptypb.Empty) (*account.AccountMeResponse, error) {
	log.Info().Str("method", "AccountMe").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing AccountMe request")

	authPayload, err := server.authenticateUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	authorized := server.authorizeUser(authPayload, [][]string{{"LSA"}})
	if !authorized {
		return nil, permissionDeniedError(fmt.Sprintln("User not authorized, missing permission: LSA"))
	}

	result, err := server.store.GetAccount(ctx, authPayload.AccountID)
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to get account:", err))
	}

	rsp := &account.AccountMeResponse{
		Account: convertAccount(result),
	}
	return rsp, nil
}
