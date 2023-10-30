package info

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *InfoServer) CreateGenres(ctx context.Context, req *nfpb.CreateGenresRequest) (*nfpb.CreateGenresResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create new genre")
	}

	if violations := validateCreateGenreRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	result, err := server.gojo.CreateGenresTx(ctx, db.CreateGenresTxParams{
		Names: req.GetNames(),
	})
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		return nil, status.Errorf(codes.Internal, "failed to create genre : %s", err)
	}

	var PBgenres []*nfpb.Genre
	for _, g := range result.Genres {
		genre := shared.ConvertGenre(g)
		PBgenres = append(PBgenres, genre)
	}

	res := &nfpb.CreateGenresResponse{
		Genres: PBgenres,
	}

	return res, nil
}

func validateCreateGenreRequest(req *nfpb.CreateGenresRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateGenreAndStudio(req.GetNames()); err != nil {
		violations = append(violations, shared.FieldViolation("names", err))
	}

	return violations
}
