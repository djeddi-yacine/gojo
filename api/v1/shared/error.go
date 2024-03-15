package shv1

import (
	"context"
	"errors"
	"fmt"
	"net"

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

func ApiError(Msg string, Err error) error {
	if errors.Is(Err, context.Canceled) {
		return status.Error(codes.Canceled, "request was canceled")
	}

	if err, ok := Err.(net.Error); ok && err.Timeout() {
		return status.Error(codes.DeadlineExceeded, "operation timed out")
	}

	statusError := status.New(codes.Internal, Msg)
	statusDetails, err := statusError.WithDetails(&errdetails.ErrorInfo{
		Reason: Err.Error(),
	})
	if err != nil {
		fmt.Println(err)
		return statusError.Err()
	}

	return statusDetails.Err()
}
