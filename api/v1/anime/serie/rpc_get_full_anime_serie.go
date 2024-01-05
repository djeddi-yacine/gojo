package asapiv1

import (
	"context"

	aapiv1 "github.com/dj-yacine-flutter/gojo/api/v1/anime"
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
		ID:      req.AnimeID,
		Target:  ping.AnimeSerie,
		Version: ping.V1,
	}

	res := &aspbv1.GetFullAnimeSerieResponse{}

	if err = server.ping.Handle(ctx, cache.Main(), &res.AnimeSerie, func() error {
		animeSerie, err := server.gojo.GetAnimeSerie(ctx, req.GetAnimeID())
		if err != nil {
			return shv1.ApiError("failed to get the anime serie", err)
		}

		res.AnimeSerie = convertAnimeSerie(animeSerie)
		return nil
	}); err != nil {
		return nil, err
	}

	if err = server.ping.Handle(ctx, cache.Meta(uint32(req.LanguageID)), &res.AnimeMeta, func() error {
		_, err := server.gojo.GetLanguage(ctx, req.GetLanguageID())
		if err != nil {
			return shv1.ApiError("failed to get the language", err)
		}

		animeMeta, err := server.gojo.GetAnimeSerieMeta(ctx, db.GetAnimeSerieMetaParams{
			AnimeID:    req.GetAnimeID(),
			LanguageID: req.GetLanguageID(),
		})
		if err != nil {
			return shv1.ApiError("no anime serie found with this language ID", err)
		}

		if animeMeta > 0 {
			meta, err := server.gojo.GetMeta(ctx, animeMeta)
			if err != nil {
				return shv1.ApiError("failed to get anime serie metadata", err)
			}

			res.AnimeMeta = &nfpbv1.AnimeMetaResponse{
				LanguageID: req.GetLanguageID(),
				Meta:       shv1.ConvertMeta(meta),
				CreatedAt:  timestamppb.New(meta.CreatedAt),
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	if err = server.ping.Handle(ctx, cache.Links(), &res.AnimeLinks, func() error {
		animeLinkID, err := server.gojo.GetAnimeSerieLink(ctx, req.GetAnimeID())
		if err != nil {
			if db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("failed to get anime serie links ID", err)
			} else {
				return nil
			}
		}

		if animeLinkID.AnimeID == req.AnimeID {
			animeLinks, err := server.gojo.GetAnimeLink(ctx, animeLinkID.LinkID)
			if err != nil {
				if db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("failed to get anime serie links", err)
				} else {
					return nil
				}
			}
			res.AnimeLinks = aapiv1.ConvertAnimeLink(animeLinks)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	if err = server.ping.Handle(ctx, cache.Images(), &res.AnimeImages, func() error {
		animePosterIDs, err := server.gojo.ListAnimeSeriePosterImages(ctx, req.AnimeID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shv1.ApiError("cannot get anime serie posters images IDs", err)
		}

		var animePosters []db.AnimeImage
		if len(animePosterIDs) > 0 {
			animePosters = make([]db.AnimeImage, len(animePosterIDs))

			for i, p := range animePosterIDs {
				poster, err := server.gojo.GetAnimeImage(ctx, p)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime serie poster image", err)
				}
				animePosters[i] = poster
			}
		}

		animeBackdropIDs, err := server.gojo.ListAnimeSerieBackdropImages(ctx, req.AnimeID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shv1.ApiError("cannot get anime serie backdrops images IDs", err)
		}

		var animeBackdrops []db.AnimeImage
		if len(animeBackdropIDs) > 0 {
			animeBackdrops = make([]db.AnimeImage, len(animeBackdropIDs))

			for i, p := range animeBackdropIDs {
				backdrop, err := server.gojo.GetAnimeImage(ctx, p)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime serie backdrop image", err)
				}
				animeBackdrops[i] = backdrop
			}
		}

		animeLogoIDs, err := server.gojo.ListAnimeSerieLogoImages(ctx, req.AnimeID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shv1.ApiError("cannot get anime serie logos images IDs", err)
		}

		var animeLogos []db.AnimeImage
		if len(animeLogoIDs) > 0 {
			animeLogos = make([]db.AnimeImage, len(animeLogoIDs))

			for i, p := range animeLogoIDs {
				logo, err := server.gojo.GetAnimeImage(ctx, p)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime serie logo image", err)
				}
				animeLogos[i] = logo
			}
		}

		res.AnimeImages = &ashpbv1.AnimeImageResponse{
			Posters:   aapiv1.ConvertAnimeImages(animePosters),
			Backdrops: aapiv1.ConvertAnimeImages(animeBackdrops),
			Logos:     aapiv1.ConvertAnimeImages(animeLogos),
		}

		return nil
	}); err != nil {
		return nil, err
	}

	if err = server.ping.Handle(ctx, cache.Trailers(), &res.AnimeTrailers, func() error {
		animeTrailerIDs, err := server.gojo.ListAnimeSerieTrailers(ctx, req.AnimeID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shv1.ApiError("cannot get anime serie trailers IDs", err)
		}

		var animeTrailers []db.AnimeTrailer
		if len(animeTrailerIDs) > 0 {
			animeTrailers = make([]db.AnimeTrailer, len(animeTrailerIDs))

			for i, t := range animeTrailerIDs {
				trailer, err := server.gojo.GetAnimeTrailer(ctx, t.TrailerID)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime serie trailer", err)
				}
				animeTrailers[i] = trailer
			}
		}

		res.AnimeTrailers = aapiv1.ConvertAnimeTrailers(animeTrailers)
		return nil
	}); err != nil {
		return nil, err
	}

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
