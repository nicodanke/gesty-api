package gapi

import (
	"context"
	"errors"
	"fmt"
	"strings"

	db "github.com/nicodanke/gesty-api/services/account-service/db/sqlc"
	"github.com/nicodanke/gesty-api/shared/proto/account-service/requests/login"
	"github.com/nicodanke/gesty-api/shared/utils"
	userValidator "github.com/nicodanke/gesty-api/services/account-service/validators/user"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"github.com/rs/zerolog/log"
)

func (server *Server) Login(ctx context.Context, req *login.LoginRequest) (*login.LoginResponse, error) {
	log.Info().Str("method", "Login").Str("request", fmt.Sprintf("%+v", req)).Msg("Processing Login request")

	violations := validateLoginRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	fields := strings.Split(req.GetUsername(), "@")
	account, err := server.store.GetAccountByCode(ctx, fields[1])
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, notFoundError(ACCOUNT_NOT_FOUND, fmt.Sprintln("Account not found:", err))
		}
		return nil, internalError(fmt.Sprintln("Failed to get account:", err))
	}

	if !account.Active {
		return nil, unprocessableError(ACCOUNT_NOT_ACTIVE, "Account not active")
	}

	user, err := server.store.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, notFoundError(USER_NOT_FOUND, fmt.Sprintln("User not found:", err))
		}
		return nil, internalError(fmt.Sprintln("Failed to get user:", err))
	}

	if !user.Active {
		return nil, unprocessableError(USER_NOT_ACTIVE, "User not active")
	}

	err = utils.CheckPassword(req.GetPassword(), user.Password)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("Invalid password:", err))
	}

	permissions, err := server.store.GetUserPermissions(ctx, user.ID)
	if err != nil && !errors.Is(err, db.ErrRecordNotFound) {
		return nil, internalError(fmt.Sprintln("Failed to get permissions:", err))
	}
	

	if user.IsAdmin {
		permissions = append([]string{"A"}, permissions...)
	}

	modules, err := server.store.GetAccountModules(ctx, account.ID)
	if err != nil && !errors.Is(err, db.ErrRecordNotFound) {
		return nil, internalError(fmt.Sprintln("Failed to get modules:", err))
	}
	
	accessToken, accessPayload, err := server.tokenMaker.CreateToken(user.ID, account.ID, account.Code, permissions, modules, server.config.AccessTokenDuration)
	if err != nil {
		return nil, internalError(fmt.Sprintln("Error creating accessToken:", err))
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(user.ID, account.ID, account.Code, permissions, modules, server.config.RefreshTokenDuration)
	if err != nil {
		return nil, internalError(fmt.Sprintln("Error creating refreshToken:", err))
	}

	metadata := server.extractMetadata(ctx)

	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		UserID:       user.ID,
		RefreshToken: refreshToken,
		UserAgent:    metadata.UserAgent,
		ClientIp:     metadata.ClientIP,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error creating session: %s", err)
	}

	rsp := &login.LoginResponse{
		User:                  convertUser(user),
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}
	return rsp, nil
}

func validateLoginRequest(req *login.LoginRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := userValidator.ValidateFullUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if err := userValidator.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	return violations
}