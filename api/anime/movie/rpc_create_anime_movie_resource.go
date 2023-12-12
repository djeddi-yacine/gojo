package animeMovie

import (
	"context"
	"errors"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/ampb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeMovieServer) CreateAnimeMovieResource(ctx context.Context, req *ampb.CreateAnimeMovieResourceRequest) (*ampb.CreateAnimeMovieResourceResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime movie resource")
	}

	if violations := validateCreateAnimeMovieResourceRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	arg := db.CreateAnimeMovieResourceTxParams{
		AnimeID: req.GetAnimeID(),
		CreateAnimeResourceParams: db.CreateAnimeResourceParams{
			TvdbID:        req.GetAnimeResources().GetTvdbID(),
			TmdbID:        req.GetAnimeResources().GetTmdbID(),
			ImdbID:        req.GetAnimeResources().GetImdbID(),
			LivechartID:   req.GetAnimeResources().GetLivechartID(),
			AnimePlanetID: req.GetAnimeResources().GetAnimePlanetID(),
			AnisearchID:   req.GetAnimeResources().GetAnisearchID(),
			AnidbID:       req.GetAnimeResources().GetAnidbID(),
			KitsuID:       req.GetAnimeResources().GetKitsuID(),
			MalID:         req.GetAnimeResources().GetMalID(),
			NotifyMoeID:   req.GetAnimeResources().GetNotifyMoeID(),
			AnilistID:     req.GetAnimeResources().GetAnilistID(),
		},
	}

	data, err := server.gojo.CreateAnimeMovieResourceTx(ctx, arg)
	if err != nil {
		return nil, shared.DatabaseError("failed to create anime movie resource", err)
	}

	res := &ampb.CreateAnimeMovieResourceResponse{
		AnimeMovie:     shared.ConvertAnimeMovie(data.AnimeMovie),
		AnimeResources: shared.ConvertAnimeResource(data.AnimeResource),
	}
	return res, nil
}

func validateCreateAnimeMovieResourceRequest(req *ampb.CreateAnimeMovieResourceRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shared.FieldViolation("ID", err))
	}

	if req.AnimeResources != nil {
		if err := utils.ValidateInt(int64(req.GetAnimeResources().GetTvdbID())); err != nil {
			violations = append(violations, shared.FieldViolation("tvdbID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeResources().GetTmdbID())); err != nil {
			violations = append(violations, shared.FieldViolation("tmdbID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeResources().GetLivechartID())); err != nil {
			violations = append(violations, shared.FieldViolation("livechartID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeResources().GetAnidbID())); err != nil {
			violations = append(violations, shared.FieldViolation("anidbID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeResources().GetAnisearchID())); err != nil {
			violations = append(violations, shared.FieldViolation("anisearchID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeResources().GetKitsuID())); err != nil {
			violations = append(violations, shared.FieldViolation("kitsuID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeResources().GetMalID())); err != nil {
			violations = append(violations, shared.FieldViolation("malID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeResources().GetAnilistID())); err != nil {
			violations = append(violations, shared.FieldViolation("anilistID", err))
		}

	} else {
		violations = append(violations, shared.FieldViolation("animeResources", errors.New("you need to send the animeResources model")))
	}

	return violations
}
