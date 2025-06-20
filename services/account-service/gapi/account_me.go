package gapi

import (
	"context"
	"fmt"

	"github.com/nicodanke/gesty-api/shared/proto/account-service/requests/account"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) AccountMe(ctx context.Context, req *emptypb.Empty) (*account.AccountMeResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
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
