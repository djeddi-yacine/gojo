package asapiv1

import (
	"context"

	aapiv1 "github.com/dj-yacine-flutter/gojo/api/v1/anime"
	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgerrcode"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeSerieServer) GetFullAnimeSeason(ctx context.Context, req *aspbv1.GetFullAnimeSeasonRequest) (*aspbv1.GetFullAnimeSeasonResponse, error) {
	var err error

	_, err = shv1.AuthorizeUser(ctx, server.tokenMaker, utils.AllRolls)
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	violations := validateGetFullAnimeSeasonRequest(req)
	if violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	cache := &ping.CacheKey{
		ID:      req.SeasonID,
		Target:  ping.AnimeSeason,
		Version: ping.V1,
	}

	res := &aspbv1.GetFullAnimeSeasonResponse{}

	if err = server.ping.Handle(ctx, cache.Main(), &res.AnimeSeason, func() error {
		animeSeason, err := server.gojo.GetAnimeSeason(ctx, req.GetSeasonID())
		if err != nil {
			return shv1.ApiError("failed to get the anime season", err)
		}

		res.AnimeSeason = convertAnimeSeason(animeSeason)
		return nil
	}); err != nil {
		return nil, err
	}

	if err = server.ping.Handle(ctx, cache.Meta(uint32(req.LanguageID)), &res.SeasonMeta, func() error {
		_, err := server.gojo.GetLanguage(ctx, req.GetLanguageID())
		if err != nil {
			return shv1.ApiError("failed to get the language", err)
		}

		animeMeta, err := server.gojo.GetAnimeSeasonMeta(ctx, db.GetAnimeSeasonMetaParams{
			SeasonID:   req.GetSeasonID(),
			LanguageID: req.GetLanguageID(),
		})
		if err != nil {
			return shv1.ApiError("no anime season found with this language ID", err)
		}

		if animeMeta > 0 {
			meta, err := server.gojo.GetMeta(ctx, animeMeta)
			if err != nil {
				return shv1.ApiError("failed to get anime season metadata", err)
			}

			res.SeasonMeta = &nfpbv1.AnimeMetaResponse{
				LanguageID: req.GetLanguageID(),
				Meta:       shv1.ConvertMeta(meta),
				CreatedAt:  timestamppb.New(meta.CreatedAt),
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	if err = server.ping.Handle(ctx, cache.Genre(), &res.SeasonGenres, func() error {
		seasonGenres, err := server.gojo.ListAnimeSeasonGenres(ctx, req.GetSeasonID())
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shv1.ApiError("failed to get anime serie genres", err)
		}

		var genres []db.Genre
		if len(seasonGenres) > 0 {
			genres = make([]db.Genre, len(seasonGenres))

			for i, amg := range seasonGenres {
				genres[i], err = server.gojo.GetGenre(ctx, amg)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("failed when list anime serie genres", err)
				}
			}
		}

		res.SeasonGenres = shv1.ConvertGenres(genres)
		return nil
	}); err != nil {
		return nil, err
	}

	if err = server.ping.Handle(ctx, cache.Studio(), &res.SeasonStudios, func() error {
		seasonStudios, err := server.gojo.ListAnimeSeasonStudios(ctx, req.GetSeasonID())
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shv1.ApiError("failed to get anime serie studios", err)
		}

		var studios []db.Studio
		if len(seasonStudios) > 0 {
			studios = make([]db.Studio, len(seasonStudios))
			for i, ams := range seasonStudios {
				studios[i], err = server.gojo.GetStudio(ctx, ams)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("failed when list anime serie studios", err)
				}
			}
		}

		res.SeasonStudios = shv1.ConvertStudios(studios)
		return nil
	}); err != nil {
		return nil, err
	}

	if err = server.ping.Handle(ctx, cache.Resources(), &res.SeasonResoures, func() error {
		seasonResourceID, err := server.gojo.GetAnimeSeasonResource(ctx, req.GetSeasonID())
		if err != nil {
			if db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot get anime season resources", err)
			} else {
				return nil
			}
		}

		seasonResources, err := server.gojo.GetAnimeResource(ctx, seasonResourceID.ResourceID)
		if err != nil {
			if db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot get resources data", err)
			} else {
				return nil
			}
		}

		res.SeasonResoures = aapiv1.ConvertAnimeResource(seasonResources)
		return nil
	}); err != nil {
		return nil, err
	}

	if err = server.ping.Handle(ctx, cache.Tags(), &res.SeasonTags, func() error {
		animeTagIDs, err := server.gojo.ListAnimeSeasonTags(ctx, req.SeasonID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shv1.ApiError("cannot get anime season tags IDs", err)
		}

		var animeTags []db.AnimeTag
		if len(animeTagIDs) > 0 {
			animeTags = make([]db.AnimeTag, len(animeTagIDs))

			for i, t := range animeTagIDs {
				tag, err := server.gojo.GetAnimeTag(ctx, t.TagID)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime season tag", err)
				}
				animeTags[i] = tag
			}
		}

		if len(animeTags) > 0 {
			res.SeasonTags = make([]*aspbv1.AnimeSeasonTag, len(animeTags))

			for i, t := range animeTags {
				res.SeasonTags[i] = &aspbv1.AnimeSeasonTag{
					ID:        t.ID,
					Tag:       t.Tag,
					CreatedAt: timestamppb.New(t.CreatedAt),
				}
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	if err = server.ping.Handle(ctx, cache.Images(), &res.SeasonPosters, func() error {
		animePosterIDs, err := server.gojo.ListAnimeSeriePosterImages(ctx, req.SeasonID)
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

		res.SeasonPosters = aapiv1.ConvertAnimeImages(animePosters)

		return nil
	}); err != nil {
		return nil, err
	}

	if err = server.ping.Handle(ctx, cache.Trailers(), &res.SeasonTrailers, func() error {
		animeTrailerIDs, err := server.gojo.ListAnimeSerieTrailers(ctx, req.SeasonID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shv1.ApiError("cannot get anime season trailers IDs", err)
		}

		var animeTrailers []db.AnimeTrailer
		if len(animeTrailerIDs) > 0 {
			animeTrailers = make([]db.AnimeTrailer, len(animeTrailerIDs))

			for i, t := range animeTrailerIDs {
				trailer, err := server.gojo.GetAnimeTrailer(ctx, t.TrailerID)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime season trailer", err)
				}
				animeTrailers[i] = trailer
			}
		}

		res.SeasonTrailers = aapiv1.ConvertAnimeTrailers(animeTrailers)
		return nil
	}); err != nil {
		return nil, err
	}

	return res, nil
}

func validateGetFullAnimeSeasonRequest(req *aspbv1.GetFullAnimeSeasonRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetSeasonID()); err != nil {
		violations = append(violations, shv1.FieldViolation("seasonID", err))
	}

	if err := utils.ValidateInt(int64(req.GetLanguageID())); err != nil {
		violations = append(violations, shv1.FieldViolation("languageID", err))
	}

	return violations
}
