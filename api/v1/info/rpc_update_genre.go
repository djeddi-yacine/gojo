package nfapiv1

import (
	"context"
	"errors"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *InfoServer) UpdateGenre(ctx context.Context, req *nfpbv1.UpdateGenreRequest) (*nfpbv1.UpdateGenreResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot update genre")
	}

	if violations := validateUpdateGenreRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	data, err := server.gojo.UpdateGenre(ctx, db.UpdateGenreParams{
		ID: req.GetGenreID(),
		GenreName: pgtype.Text{
			String: req.GetGenreName(),
			Valid:  req.GenreName != nil,
		},
	})
	if err != nil {
		return nil, shv1.ApiError("failed to update genre", err)
	}

	res := &nfpbv1.UpdateGenreResponse{
		Genre: shv1.ConvertGenre(data),
	}

	return res, nil
}

func validateUpdateGenreRequest(req *nfpbv1.UpdateGenreRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(int64(req.GetGenreID())); err != nil {
		violations = append(violations, shv1.FieldViolation("genreID", err))
	}

	if req.GenreName != nil {
		if err := utils.ValidateString(req.GetGenreName(), 2, 15); err != nil {
			violations = append(violations, shv1.FieldViolation("genreName", err))
		}
	} else {
		violations = append(violations, shv1.FieldViolation("genreName", errors.New("put the genre name")))
	}

	return violations
}
