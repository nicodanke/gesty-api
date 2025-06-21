package gapi

import (
	"context"
	"errors"
	"fmt"
	"time"

	db "github.com/nicodanke/gesty-api/services/account-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/account-service/requests/login"
	"github.com/nicodanke/gesty-api/services/account-service/validators"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"github.com/rs/zerolog/log"
)

func (server *Server) RefreshToken(ctx context.Context, req *login.RefreshTokenRequest) (*login.RefreshTokenResponse, error) {
	log.Info().Str("method", "RefreshToken").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing RefreshToken request")

	violations := validateRefreshTokenRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	refreshPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid refresh token: %s", err)
	}

	session, err := server.store.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "Session does not exist: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "Failed to get session: %s", err)
	}

	if session.IsBlocked {
		err := fmt.Errorf("blocked Session")
		return nil, status.Errorf(codes.Unauthenticated, "Error: %s", err)
	}

	if session.UserID != refreshPayload.UserID {
		err := fmt.Errorf("incorrect session user")
		return nil, status.Errorf(codes.Unauthenticated, "Error: %s", err)
	}

	if session.RefreshToken != req.RefreshToken {
		err := fmt.Errorf("mismatched session token")
		return nil, status.Errorf(codes.Unauthenticated, "Error: %s", err)
	}

	if time.Now().After(session.ExpiresAt) {
		err := fmt.Errorf("expired session")
		return nil, status.Errorf(codes.Unauthenticated, "Error: %s", err)
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(refreshPayload.UserID, refreshPayload.AccountID, refreshPayload.AccountCode, refreshPayload.Permissions, refreshPayload.Modules, server.config.AccessTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %s", err)
	}

	rsp := &login.RefreshTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: timestamppb.New(accessPayload.ExpiredAt),
	}
	return rsp, nil
}

func validateRefreshTokenRequest(req *login.RefreshTokenRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validators.ValidString(req.GetRefreshToken(), 10, 10000); err != nil {
		violations = append(violations, fieldViolation("refreshToken", err))
	}

	return violations
}