package asapiv1

import (
	"context"

	av1 "github.com/dj-yacine-flutter/gojo/api/v1/anime"
	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ashpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ashpb"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgerrcode"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeSerieServer) GetFullAnimeSerie(ctx context.Context, req *aspbv1.GetFullAnimeSerieRequest) (*aspbv1.GetFullAnimeSerieResponse, error) {
	var err error

	_, err = shv1.AuthorizeUser(ctx, server.tokenMaker, utils.AllRolls)
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	violations := validateGetFullAnimeSerieRequest(req)
	if violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	cache := &ping.CacheKey{
		ID:     req.GetAnimeID(),
		Target: ping.AnimeSerie,
	}

	res := &aspbv1.GetFullAnimeSerieResponse{}

	var serie db.AnimeSerie
	if err = server.ping.Handle(ctx, cache.Main(), &serie, func() error {
		serie, err = server.gojo.GetAnimeSerie(ctx, req.GetAnimeID())
		if err != nil {
			return shv1.ApiError("cannot get anime serie", err)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	res.AnimeSerie = convertAnimeSerie(serie)

	var meta db.Meta
	if err = server.ping.Handle(ctx, cache.Meta(), &meta, func() error {
		animeMeta, err := server.gojo.GetAnimeSerieMeta(ctx, db.GetAnimeSerieMetaParams{
			AnimeID:    req.GetAnimeID(),
			LanguageID: req.GetLanguageID(),
		})
		if err != nil {
			return shv1.ApiError("no anime serie found with this language ID", err)
		}

		if animeMeta > 0 {
			meta, err = server.gojo.GetMeta(ctx, animeMeta)
			if err != nil {
				return shv1.ApiError("cannot get anime serie metadata", err)
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	res.AnimeMeta = &nfpbv1.AnimeMetaResponse{
		LanguageID: req.GetLanguageID(),
		Meta:       shv1.ConvertMeta(meta),
		CreatedAt:  timestamppb.New(meta.CreatedAt),
	}

	var link db.AnimeLink
	if err = server.ping.Handle(ctx, cache.Links(), &link, func() error {
		ID, err := server.gojo.GetAnimeSerieLink(ctx, req.GetAnimeID())
		if err != nil {
			if db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot get anime serie links ID", err)
			} else {
				return nil
			}
		}

		link, err = server.gojo.GetAnimeLink(ctx, ID.ID)
		if err != nil {
			if db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot get anime serie links", err)
			} else {
				return nil
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	res.AnimeLinks = av1.ConvertAnimeLink(link)

	var pIDs []int64
	if err = server.ping.Handle(ctx, cache.Posters(), &pIDs, func() error {
		pIDs, err = server.gojo.ListAnimeSeriePosterImages(ctx, req.GetAnimeID())
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shv1.ApiError("cannot get anime serie posters images IDs", err)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	posters, err := server.gojo.ListAnimeImagesTx(ctx, pIDs)
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shv1.ApiError("cannot get anime serie posters images", err)
	}

	var bIDs []int64
	if err = server.ping.Handle(ctx, cache.Backdrops(), &bIDs, func() error {
		bIDs, err = server.gojo.ListAnimeSerieBackdropImages(ctx, req.GetAnimeID())
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shv1.ApiError("cannot get anime serie backdrops images IDs", err)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	backdrops, err := server.gojo.ListAnimeImagesTx(ctx, bIDs)
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shv1.ApiError("cannot get anime serie backdrops images", err)
	}

	var lIDs []int64
	if err = server.ping.Handle(ctx, cache.Logos(), &lIDs, func() error {
		lIDs, err = server.gojo.ListAnimeSerieLogoImages(ctx, req.GetAnimeID())
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shv1.ApiError("cannot get anime serie logos images IDs", err)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	logos, err := server.gojo.ListAnimeImagesTx(ctx, lIDs)
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shv1.ApiError("cannot get anime serie logos images", err)
	}

	res.AnimeImages = &ashpbv1.AnimeImageResponse{
		Posters:   av1.ConvertAnimeImages(posters),
		Backdrops: av1.ConvertAnimeImages(backdrops),
		Logos:     av1.ConvertAnimeImages(logos),
	}

	var rIDs []int64
	if err = server.ping.Handle(ctx, cache.Trailers(), &rIDs, func() error {
		rIDs, err = server.gojo.ListAnimeSerieTrailers(ctx, req.GetAnimeID())
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shv1.ApiError("cannot get anime serie trailers IDs", err)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	trailers, err := server.gojo.ListAnimeTrailersTx(ctx, rIDs)
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shv1.ApiError("cannot get anime serie trailers", err)
	}

	res.AnimeTrailers = av1.ConvertAnimeTrailers(trailers)

	return res, nil
}

func validateGetFullAnimeSerieRequest(req *aspbv1.GetFullAnimeSerieRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shv1.FieldViolation("animeID", err))
	}

	if err := utils.ValidateInt(int64(req.GetLanguageID())); err != nil {
		violations = append(violations, shv1.FieldViolation("languageID", err))
	}

	return violations
}
