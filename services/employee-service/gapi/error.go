package gapi

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	INTERNAL           		string = "INTERNAL"
	UNAUTHENTICATED    		string = "UNAUTHENTICATED"
	FORBIDDEN          		string = "FORBIDDEN"
	UNPROCESSABLE_ENTITY    string = "UNPROCESSABLE_ENTITY"
	CONFLICT                string = "CONFLICT"
	CONFLICT_FK             string = "CONFLICT_FK"
	CONFLICT_UNIQUE         string = "CONFLICT_UNIQUE"
	
	ACTION_NOT_ALLOWED 		string = "ACTION_NOT_ALLOWED"

	ACCOUNT_NOT_ACTIVE 		string = "ACCOUNT_NOT_ACTIVE"
	USER_NOT_ACTIVE    		string = "USER_NOT_ACTIVE"
	ACTION_NOT_DELETABLE    string = "ACTION_NOT_DELETABLE"

	NOT_FOUND	            string = "NOT_FOUND"
	ACCOUNT_NOT_FOUND  		string = "ACCOUNT_NOT_FOUND"
	USER_NOT_FOUND     		string = "USER_NOT_FOUND"
	ROLE_NOT_FOUND     		string = "ROLE_NOT_FOUND"
)

func fieldViolation(field string, err error) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: err.Error(),
	}
}

// This error returns a 400 status code
func invalidArgumentError(violations []*errdetails.BadRequest_FieldViolation) error {
	badRequest := &errdetails.BadRequest{FieldViolations: violations}
	statusInvalid := status.New(codes.InvalidArgument, "Invalid parameters")
	statusDetails, err := statusInvalid.WithDetails(badRequest)
	if err != nil {
		return statusInvalid.Err()
	}
	return statusDetails.Err()
}

// This error returns a 404 status code
func notFoundError(msg string, reason string) error {
	if msg == "" {
		msg = NOT_FOUND
	}
	return GetError(codes.NotFound, msg, reason)
}

// This error returns a 401 status code
func unauthenticatedError(msg string) error {
	return GetError(codes.Unauthenticated, UNAUTHENTICATED, msg)
}

// This error returns a 500 status code
func internalError(reason string) error {
	return GetError(codes.Internal, INTERNAL, reason)
}

// This error returns a 422 status code
func unprocessableError(msg string, reason string) error {
	if msg == "" {
		msg = UNPROCESSABLE_ENTITY
	}
	return GetError(codes.FailedPrecondition, msg, reason)
}

// This error returns a 409 status code
func conflictError(msg string, reason string, constraintName string) error {
	if msg == "" {
		msg = CONFLICT
	}
	return GetErrorConstraint(codes.AlreadyExists, msg, reason, constraintName)
}

// This error returns a 403 status code
func permissionDeniedError(reason string) error {
	return GetError(codes.PermissionDenied, FORBIDDEN, reason)
}

func GetError(code codes.Code, msg string, reason string) error {
	st := status.New(code, msg)

	if reason != "" {
		ds, err := st.WithDetails(
			&errdetails.ErrorInfo{
				Reason: reason,
			},
		)
		if err != nil {
			return st.Err()
		}
		return ds.Err()
	}

	return st.Err()
}

func GetErrorConstraint(code codes.Code, msg string, reason string, constraintName string) error {
	st := status.New(code, msg)

	if reason != "" {
		ds, err := st.WithDetails(
			&errdetails.ErrorInfo{
				Reason: reason,
				Metadata: map[string]string{
					"constraint_name": constraintName,
				},
			},
		)
		if err != nil {
			return st.Err()
		}
		return ds.Err()
	}

	return st.Err()
}
