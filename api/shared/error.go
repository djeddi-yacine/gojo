package shared

import (
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

func DatabaseError(msg string, err error) error {
	dberr := db.ErrorDB(err)
	errdet := &errdetails.ErrorInfo{
		Reason: msg,
	}
	if dberr != nil {
		errdet.Metadata = map[string]string{
			"Statue":        db.ErrorDBType(err),
			"Database Code": dberr.Code,
			//			"Database Message": dberr.Message,
			"Database Details": dberr.Detail,
		}
	}

	var statusError *status.Status
	if dberr.Code == pgerrcode.CaseNotFound {
		statusError = status.New(codes.NotFound, "Internal server")
	} else if dberr.Code == pgerrcode.UniqueViolation {
		statusError = status.New(codes.AlreadyExists, "Internal server")
	} else {
		statusError = status.New(codes.Internal, "Internal server")
	}

	statusDetails, err := statusError.WithDetails(errdet)
	if err != nil {
		return statusError.Err()
	}

	return statusDetails.Err()
}
