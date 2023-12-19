package animeSerie

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/aspb"
	"github.com/dj-yacine-flutter/gojo/pb/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgerrcode"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeSerieServer) GetFullAnimeSeason(ctx context.Context, req *aspb.GetFullAnimeSeasonRequest) (*aspb.GetFullAnimeSeasonResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot get full anime serie")
	}

	violations := validateGetFullAnimeSeasonRequest(req)
	if violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	animeSeason, err := server.gojo.GetAnimeSeason(ctx, req.GetSeasonID())
	if err != nil {
		if db.ErrorDB(err).Code == pgerrcode.CaseNotFound {
			return nil, shared.DatabaseError("there is no anime season with this ID", err)
		}
		return nil, shared.DatabaseError("failed to get the anime season", err)
	}

	_, err = server.gojo.GetLanguage(ctx, req.GetLanguageID())
	if err != nil {
		if db.ErrorDB(err).Code == pgerrcode.CaseNotFound {
			return nil, shared.DatabaseError("there is no language with this language ID", err)
		}
		return nil, shared.DatabaseError("failed to get the language", err)
	}

	res := &aspb.GetFullAnimeSeasonResponse{
		AnimeSeason: shared.ConvertAnimeSeason(animeSeason),
	}

	animeMeta, err := server.gojo.GetAnimeSeasonMeta(ctx, db.GetAnimeSeasonMetaParams{
		SeasonID:   req.GetSeasonID(),
		LanguageID: req.GetLanguageID(),
	})
	if err != nil {
		return nil, shared.DatabaseError("no anime season found with this language ID", err)
	}

	if animeMeta > 0 {
		meta, err := server.gojo.GetMeta(ctx, animeMeta)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return nil, shared.DatabaseError("failed to get anime season metadata", err)
		}

		res.SeasonMeta = &nfpb.AnimeMetaResponse{
			LanguageID: req.GetLanguageID(),
			Meta:       shared.ConvertMeta(meta),
			CreatedAt:  timestamppb.New(meta.CreatedAt),
		}
	}

	seasonGenres, err := server.gojo.ListAnimeSeasonGenres(ctx, req.GetSeasonID())
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shared.DatabaseError("failed to get anime serie genres", err)
	}

	if len(seasonGenres) > 0 {
		genres := make([]db.Genre, len(seasonGenres))

		for i, amg := range seasonGenres {
			genres[i], err = server.gojo.GetGenre(ctx, amg)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return nil, shared.DatabaseError("failed when list anime serie genres", err)
			}
		}
		res.SeasonGenres = shared.ConvertGenres(genres)
	}

	seasonStudios, err := server.gojo.ListAnimeSeasonStudios(ctx, req.GetSeasonID())
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shared.DatabaseError("failed to get anime serie studios", err)
	}

	if len(seasonStudios) > 0 {
		studios := make([]db.Studio, len(seasonStudios))
		for i, ams := range seasonStudios {
			studios[i], err = server.gojo.GetStudio(ctx, ams)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return nil, shared.DatabaseError("failed when list anime serie studios", err)
			}
		}
		res.SeasonStudios = shared.ConvertStudios(studios)
	}

	seasonResourceID, err := server.gojo.GetAnimeSeasonResource(ctx, req.GetSeasonID())
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shared.DatabaseError("cannot get anime season resources", err)
	}

	seasonResources, err := server.gojo.GetAnimeResource(ctx, seasonResourceID.ResourceID)
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shared.DatabaseError("cannot get resources data", err)
	}
	res.SeasonResoures = shared.ConvertAnimeResource(seasonResources)

	animeTagIDs, err := server.gojo.ListAnimeSeasonTags(ctx, req.SeasonID)
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shared.DatabaseError("cannot get anime season tags IDs", err)
	}

	var animeTags []db.AnimeTag
	if len(animeTagIDs) > 0 {
		animeTags = make([]db.AnimeTag, len(animeTagIDs))

		for i, t := range animeTagIDs {
			tag, err := server.gojo.GetAnimeTag(ctx, t.TagID)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return nil, shared.DatabaseError("cannot get anime season tag", err)
			}
			animeTags[i] = tag
		}
	}

	if len(animeTags) > 0 {
		res.SeasonTags = make([]*aspb.AnimeSeasonTag, len(animeTags))

		for i, t := range animeTags {
			res.SeasonTags[i] = &aspb.AnimeSeasonTag{
				ID:        t.ID,
				Tag:       t.Tag,
				CreatedAt: timestamppb.New(t.CreatedAt),
			}
		}
	}

	animePosterIDs, err := server.gojo.ListAnimeSeriePosterImages(ctx, req.SeasonID)
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shared.DatabaseError("cannot get anime serie posters images IDs", err)
	}

	var animePosters []db.AnimeImage
	if len(animePosterIDs) > 0 {
		animePosters = make([]db.AnimeImage, len(animePosterIDs))

		for i, p := range animePosterIDs {
			poster, err := server.gojo.GetAnimeImage(ctx, p)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return nil, shared.DatabaseError("cannot get anime serie poster image", err)
			}
			animePosters[i] = poster
		}
	}

	animeBackdropIDs, err := server.gojo.ListAnimeSerieBackdropImages(ctx, req.SeasonID)
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shared.DatabaseError("cannot get anime serie backdrops images IDs", err)
	}

	var animeBackdrops []db.AnimeImage
	if len(animeBackdropIDs) > 0 {
		animeBackdrops = make([]db.AnimeImage, len(animeBackdropIDs))

		for i, p := range animeBackdropIDs {
			backdrop, err := server.gojo.GetAnimeImage(ctx, p)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return nil, shared.DatabaseError("cannot get anime serie backdrop image", err)
			}
			animeBackdrops[i] = backdrop
		}
	}

	animeLogoIDs, err := server.gojo.ListAnimeSerieLogoImages(ctx, req.SeasonID)
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shared.DatabaseError("cannot get anime serie logos images IDs", err)
	}

	var animeLogos []db.AnimeImage
	if len(animeLogoIDs) > 0 {
		animeLogos = make([]db.AnimeImage, len(animeLogoIDs))

		for i, p := range animeLogoIDs {
			logo, err := server.gojo.GetAnimeImage(ctx, p)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return nil, shared.DatabaseError("cannot get anime serie logo image", err)
			}
			animeLogos[i] = logo
		}
	}

	res.SeasonPosters = shared.ConvertAnimeImages(animePosters)

	animeTrailerIDs, err := server.gojo.ListAnimeSerieTrailers(ctx, req.SeasonID)
	if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
		return nil, shared.DatabaseError("cannot get anime season trailers IDs", err)
	}

	var animeTrailers []db.AnimeTrailer
	if len(animeTrailerIDs) > 0 {
		animeTrailers = make([]db.AnimeTrailer, len(animeTrailerIDs))

		for i, t := range animeTrailerIDs {
			trailer, err := server.gojo.GetAnimeTrailer(ctx, t.TrailerID)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return nil, shared.DatabaseError("cannot get anime season trailer", err)
			}
			animeTrailers[i] = trailer
		}
	}

	res.SeasonTrailers = shared.ConvertAnimeTrailers(animeTrailers)

	return res, nil
}

func validateGetFullAnimeSeasonRequest(req *aspb.GetFullAnimeSeasonRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetSeasonID()); err != nil {
		violations = append(violations, shared.FieldViolation("seasonID", err))
	}

	if err := utils.ValidateInt(int64(req.GetLanguageID())); err != nil {
		violations = append(violations, shared.FieldViolation("languageID", err))
	}

	return violations
}
