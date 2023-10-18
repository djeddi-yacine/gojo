package api

import (
	"context"

	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateAnimeMovieMeta(ctx context.Context, req *pb.CreateAnimeMovieMetaRequest) (*pb.CreateAnimeMovieMetaResponse, error) {
	authPayload, err := server.authorizeUser(ctx, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, unAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime movie metadata")
	}

	if violations := validateCreateAnimeMovieMetaRequest(req); violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.CreateAnimeMovieMetaTxParams{
		AnimeID:    req.GetAnimeID(),
		LanguageID: req.GetAnimeMeta().GetLanguageID(),
		CreateMetaParams: db.CreateMetaParams{
			Title:    req.GetAnimeMeta().GetMeta().GetTitle(),
			Overview: req.GetAnimeMeta().GetMeta().GetOverview(),
		},
	}

	anime, err := server.gojo.CreateAnimeMovieMetaTx(ctx, arg)
	if err != nil {
		db.ErrorSQL(err)
		return nil, status.Errorf(codes.Internal, "failed to create anime movie metadata : %s", err)
	}

	res := &pb.CreateAnimeMovieMetaResponse{
		AnimeID: req.GetAnimeID(),
		AnimeMeta: &pb.AnimeMetaResponse{
			Meta:      ConvertMeta(anime.Meta),
			Language:  ConvertLanguage(anime.Language),
			CreatedAt: timestamppb.New(anime.Meta.CreatedAt),
		},
	}
	return res, nil
}

func validateCreateAnimeMovieMetaRequest(req *pb.CreateAnimeMovieMetaRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, fieldViolation("animeID", err))
	}

	if err := utils.ValidateInt(int64(req.GetAnimeMeta().GetLanguageID())); err != nil {
		violations = append(violations, fieldViolation("languageID", err))
	}

	if err := utils.ValidateString(req.GetAnimeMeta().GetMeta().GetTitle(), 2, 500); err != nil {
		violations = append(violations, fieldViolation("title", err))
	}

	if err := utils.ValidateString(req.GetAnimeMeta().GetMeta().GetOverview(), 5, 5000); err != nil {
		violations = append(violations, fieldViolation("overview", err))
	}

	return violations
}
