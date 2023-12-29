package shv1

import (
	"context"
	"errors"
	"net"

	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/jackc/pgerrcode"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func FieldViolation(field string, err error) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: err.Error(),
	}
}

func InvalidArgumentError(violations []*errdetails.BadRequest_FieldViolation) error {
	badRequest := &errdetails.BadRequest{FieldViolations: violations}
	statusInvalid := status.New(codes.InvalidArgument, "invalid parameters")

	statusDetails, err := statusInvalid.WithDetails(badRequest)
	if err != nil {
		return statusInvalid.Err()
	}

	return statusDetails.Err()
}

func UnAuthenticatedError(err error) error {
	return status.Errorf(codes.Unauthenticated, "unauthorized: %s", err)
}

func ApiError(msg string, err error) error {
	if errors.Is(err, context.Canceled) {
		return status.Error(codes.Canceled, "request was canceled")
	}

	if err, ok := err.(net.Error); ok && err.Timeout() {
		return status.Error(codes.DeadlineExceeded, "operation timed out")
	}

	dberr := db.ErrorDB(err)
	errorDetails := &errdetails.ErrorInfo{
		Reason: msg,
	}

	var statusError *status.Status
	if dberr != nil {
		errorDetails.Metadata = map[string]string{
			"Statue":           db.ErrorType(err),
			"Database Code":    dberr.Code,
			"Database Message": dberr.Message,
			"Database Details": dberr.Detail,
		}
		switch dberr.Code {
		case pgerrcode.CaseNotFound:
			statusError = status.New(codes.NotFound, "internal server")
		case pgerrcode.UniqueViolation:
			statusError = status.New(codes.AlreadyExists, "internal server")
		case pgerrcode.ForeignKeyViolation:
			statusError = status.New(codes.FailedPrecondition, "internal server")
		default:
			statusError = status.New(codes.Internal, "internal server")
		}
	} else {
		statusError = status.New(codes.Internal, "internal server")
	}

	statusDetails, err := statusError.WithDetails(errorDetails)
	if err != nil {
		return statusError.Err()
	}

	return statusDetails.Err()
}
