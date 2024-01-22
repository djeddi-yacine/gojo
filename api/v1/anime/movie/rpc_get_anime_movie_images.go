package amapiv1

import (
	"context"

	av1 "github.com/dj-yacine-flutter/gojo/api/v1/anime"
	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	apbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/apb"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgerrcode"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *AnimeMovieServer) GetAnimeMovieImages(ctx context.Context, req *ampbv1.GetAnimeMovieImagesRequest) (*ampbv1.GetAnimeMovieImagesResponse, error) {
	var err error

	_, err = shv1.AuthorizeUser(ctx, server.tokenMaker, utils.AllRolls)
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	violations := validateGetAnimeMovieImagesRequest(req)
	if violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	cache := &ping.CacheKey{
		ID:     req.AnimeID,
		Target: ping.AnimeMovie,
	}

	res := &ampbv1.GetAnimeMovieImagesResponse{}

	var pIDs []int64
	if err = server.ping.Handle(ctx, cache.Posters(), &pIDs, func() error {
		pIDs, err = server.gojo.ListAnimeMoviePosterImages(ctx, req.AnimeID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shv1.ApiError("cannot get anime movie posters images IDs", err)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	posters, err := server.gojo.ListAnimeImagesTx(ctx, pIDs)
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shv1.ApiError("cannot get anime movie posters images", err)
	}

	var bIDs []int64
	if err = server.ping.Handle(ctx, cache.Backdrops(), &bIDs, func() error {
		bIDs, err = server.gojo.ListAnimeMovieBackdropImages(ctx, req.AnimeID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shv1.ApiError("cannot get anime movie backdrops images IDs", err)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	backdrops, err := server.gojo.ListAnimeImagesTx(ctx, bIDs)
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shv1.ApiError("cannot get anime movie backdrops images", err)
	}

	var lIDs []int64
	if err = server.ping.Handle(ctx, cache.Logos(), &lIDs, func() error {
		lIDs, err = server.gojo.ListAnimeMovieLogoImages(ctx, req.AnimeID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shv1.ApiError("cannot get anime movie logos images IDs", err)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	logos, err := server.gojo.ListAnimeImagesTx(ctx, lIDs)
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shv1.ApiError("cannot get anime movie logos images", err)
	}

	res.AnimeImages = &apbv1.AnimeImageResponse{
		Posters:   av1.ConvertAnimeImages(posters),
		Backdrops: av1.ConvertAnimeImages(backdrops),
		Logos:     av1.ConvertAnimeImages(logos),
	}

	return res, nil
}

func validateGetAnimeMovieImagesRequest(req *ampbv1.GetAnimeMovieImagesRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shv1.FieldViolation("animeID", err))
	}

	return violations
}
