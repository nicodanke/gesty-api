package gapi

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	ACCOUNT_NOT_ACTIVE string = "ACCOUNT_NOT_ACTIVE"
	ACCOUNT_NOT_FOUND  string = "ACCOUNT_NOT_FOUND"
	INTERNAL           string = "INTERNAL"
	UNAUTHENTICATED    string = "UNAUTHENTICATED"
	USER_NOT_ACTIVE    string = "USER_NOT_ACTIVE"
	USER_NOT_FOUND     string = "USER_NOT_FOUND"
	ACTION_NOT_ALLOWED string = "ACTION_NOT_ALLOWED"
)

func fieldViolation(field string, err error) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: err.Error(),
	}
}

func invalidArgumentError(violations []*errdetails.BadRequest_FieldViolation) error {
	badRequest := &errdetails.BadRequest{FieldViolations: violations}
	statusInvalid := status.New(codes.InvalidArgument, "invalid parameters")
	statusDetails, err := statusInvalid.WithDetails(badRequest)
	if err != nil {
		return statusInvalid.Err()
	}
	return statusDetails.Err()
}

func notFoundError(code string, msg string) error {
	return GetError(codes.NotFound, code, msg)
}

func unauthenticatedError(msg string) error {
	return GetError(codes.Unauthenticated, UNAUTHENTICATED, msg)
}

func internalError(msg string) error {
	return GetError(codes.Internal, INTERNAL, msg)
}

func unprocessableError(code string, msg string) error {
	return GetError(codes.Code(422), code, msg)
}

func permissionDeniedError(code string, msg string) error {
	return GetError(codes.PermissionDenied, code, msg)
}

func GetError(code codes.Code, codeStr string, msg string) error {
	st := status.New(code, msg)
	ds, err := st.WithDetails(
		&errdetails.ErrorInfo{
			Reason: codeStr,
		},
	)
	if err != nil {
		return st.Err()
	}
	return ds.Err()
}
