package amapiv1

import (
	"context"

	aapiv1 "github.com/dj-yacine-flutter/gojo/api/v1/anime"
	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	ashpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ashpb"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgerrcode"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeMovieServer) GetFullAnimeMovie(ctx context.Context, req *ampbv1.GetFullAnimeMovieRequest) (*ampbv1.GetFullAnimeMovieResponse, error) {
	var err error

	_, err = shv1.AuthorizeUser(ctx, server.tokenMaker, utils.AllRolls)
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	violations := validateGetFullAnimeMovieRequest(req)
	if violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	cache := &ping.CacheKey{
		ID:      req.AnimeID,
		Target:  ping.AnimeMovie,
		Version: ping.V1,
	}

	res := &ampbv1.GetFullAnimeMovieResponse{}

	if err = server.ping.Handle(ctx, cache.Main(), &res.AnimeMovie, func() error {
		animeMovie, err := server.gojo.GetAnimeMovie(ctx, req.GetAnimeID())
		if err != nil {
			return shv1.ApiError("cannot get anime movie", err)
		}

		res.AnimeMovie = convertAnimeMovie(animeMovie)
		return nil
	}); err != nil {
		return nil, err
	}

	if err = server.ping.Handle(ctx, cache.Meta(uint32(req.LanguageID)), &res.AnimeMeta, func() error {
		_, err := server.gojo.GetLanguage(ctx, req.GetLanguageID())
		if err != nil {
			return shv1.ApiError("no language found with this language ID", err)
		}

		animeMeta, err := server.gojo.GetAnimeMovieMeta(ctx, db.GetAnimeMovieMetaParams{
			AnimeID:    req.GetAnimeID(),
			LanguageID: req.GetLanguageID(),
		})
		if err != nil {
			return shv1.ApiError("no anime movie found with this language ID", err)
		}

		if animeMeta > 0 {
			meta, err := server.gojo.GetMeta(ctx, animeMeta)
			if err != nil {
				return shv1.ApiError("cannot get anime movie metadata", err)
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

	if err = server.ping.Handle(ctx, cache.Resources(), &res.AnimeResources, func() error {
		animeResourceID, err := server.gojo.GetAnimeMovieResource(ctx, req.GetAnimeID())
		if err != nil {
			if db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot get anime movie resources", err)
			} else {
				return nil
			}
		}

		animeResources, err := server.gojo.GetAnimeResource(ctx, animeResourceID.ResourceID)
		if err != nil {
			if db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot get resources data", err)
			} else {
				return nil
			}
		}

		res.AnimeResources = aapiv1.ConvertAnimeResource(animeResources)
		return nil
	}); err != nil {
		return nil, err
	}

	if err = server.ping.Handle(ctx, cache.Genre(), &res.AnimeGenres, func() error {
		animeMovieGenres, err := server.gojo.ListAnimeMovieGenres(ctx, req.GetAnimeID())
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shv1.ApiError("cannot get anime movie genres", err)
		}

		var genres []db.Genre
		if len(animeMovieGenres) > 0 {
			genres = make([]db.Genre, len(animeMovieGenres))

			for i, amg := range animeMovieGenres {
				genres[i], err = server.gojo.GetGenre(ctx, amg)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot list anime movie genres", err)
				}
			}
		}

		res.AnimeGenres = shv1.ConvertGenres(genres)
		return nil
	}); err != nil {
		return nil, err
	}

	if err = server.ping.Handle(ctx, cache.Studio(), &res.AnimeStudios, func() error {
		animeMovieStudios, err := server.gojo.ListAnimeMovieStudios(ctx, req.GetAnimeID())
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shv1.ApiError("cannot get anime movie studios", err)
		}

		var studios []db.Studio
		if len(animeMovieStudios) > 0 {
			studios = make([]db.Studio, len(animeMovieStudios))
			for i, ams := range animeMovieStudios {
				studios[i], err = server.gojo.GetStudio(ctx, ams)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot list anime movie studios", err)
				}
			}
		}

		res.AnimeStudios = shv1.ConvertStudios(studios)
		return nil
	}); err != nil {
		return nil, err
	}

	if err = server.ping.Handle(ctx, cache.Server(), &res.ServerID, func() error {
		sv, err := server.gojo.GetAnimeMovieServer(ctx, req.GetAnimeID())
		if err != nil {
			if db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot get anime movie server ID", err)
			} else {
				return nil
			}
		}

		res.ServerID = sv.ID
		return nil
	}); err != nil {
		return nil, err
	}

	if res.ServerID != 0 {
		if err = server.ping.Handle(ctx, cache.Sub(), &res.Sub, func() error {
			ss, err := server.gojo.ListAnimeMovieServerSubVideos(ctx, res.ServerID)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot list anime movie server sub videos", err)
			}

			subV := make([]db.AnimeMovieVideo, len(ss))
			for i, ks := range ss {
				subV[i], err = server.gojo.GetAnimeMovieVideo(ctx, ks.VideoID)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime movie server sub videos", err)
				}
			}

			st, err := server.gojo.ListAnimeMovieServerSubTorrents(ctx, res.ServerID)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot list anime movie server sub torrents", err)
			}

			subT := make([]db.AnimeMovieTorrent, len(st))
			for i, kst := range st {
				subT[i], err = server.gojo.GetAnimeMovieTorrent(ctx, kst.ServerID)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime movie server sub torrents", err)
				}
			}

			res.Sub = &ampbv1.AnimeMovieSubDataResponse{
				Videos:   convertAnimeMovieVideos(subV),
				Torrents: convertAnimeMovieTorrents(subT),
			}

			return nil
		}); err != nil {
			return nil, err
		}

		if err = server.ping.Handle(ctx, cache.Dub(), &res.Dub, func() error {
			sd, err := server.gojo.ListAnimeMovieServerDubVideos(ctx, res.ServerID)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot list anime movie server dub videos", err)
			}

			dubV := make([]db.AnimeMovieVideo, len(sd))
			for i, kd := range sd {
				dubV[i], err = server.gojo.GetAnimeMovieVideo(ctx, kd.VideoID)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime movie server dub videos", err)
				}
			}

			dt, err := server.gojo.ListAnimeMovieServerDubTorrents(ctx, res.ServerID)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot list anime movie server dub torrents", err)
			}

			dubT := make([]db.AnimeMovieTorrent, len(dt))
			for i, kdt := range dt {
				dubT[i], err = server.gojo.GetAnimeMovieTorrent(ctx, kdt.ServerID)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime movie server dub torrents", err)
				}
			}

			res.Dub = &ampbv1.AnimeMovieDubDataResponse{
				Videos:   convertAnimeMovieVideos(dubV),
				Torrents: convertAnimeMovieTorrents(dubT),
			}
			return nil
		}); err != nil {
			return nil, err
		}
	}

	if err = server.ping.Handle(ctx, cache.Links(), &res.AnimeLinks, func() error {
		animeLinkID, err := server.gojo.GetAnimeMovieLink(ctx, req.AnimeID)
		if err != nil {
			if db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot get anime movie links ID", err)
			} else {
				return nil
			}
		}

		if animeLinkID.AnimeID == req.AnimeID {
			animeLink, err := server.gojo.GetAnimeLink(ctx, animeLinkID.ID)
			if err != nil {
				if db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime movie links", err)
				} else {
					return nil
				}
			}

			res.AnimeLinks = aapiv1.ConvertAnimeLink(animeLink)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	if err = server.ping.Handle(ctx, cache.Tags(), &res.AnimeTags, func() error {
		animeTagIDs, err := server.gojo.ListAnimeMovieTags(ctx, req.AnimeID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shv1.ApiError("cannot get anime movie tags IDs", err)
		}

		var animeTags []db.AnimeTag
		if len(animeTagIDs) > 0 {
			animeTags = make([]db.AnimeTag, len(animeTagIDs))

			for i, t := range animeTagIDs {
				tag, err := server.gojo.GetAnimeTag(ctx, t.TagID)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime movie tag", err)
				}
				animeTags[i] = tag
			}
		}

		if len(animeTags) > 0 {
			res.AnimeTags = make([]*ampbv1.AnimeMovieTag, len(animeTags))

			for i, t := range animeTags {
				res.AnimeTags[i] = &ampbv1.AnimeMovieTag{
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

	if err = server.ping.Handle(ctx, cache.Images(), &res.AnimeImages, func() error {
		animePosterIDs, err := server.gojo.ListAnimeMoviePosterImages(ctx, req.AnimeID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shv1.ApiError("cannot get anime movie posters images IDs", err)
		}

		var animePosters []db.AnimeImage
		if len(animePosterIDs) > 0 {
			animePosters = make([]db.AnimeImage, len(animePosterIDs))

			for i, p := range animePosterIDs {
				poster, err := server.gojo.GetAnimeImage(ctx, p)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime movie poster image", err)
				}
				animePosters[i] = poster
			}
		}

		animeBackdropIDs, err := server.gojo.ListAnimeMovieBackdropImages(ctx, req.AnimeID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shv1.ApiError("cannot get anime movie backdrops images IDs", err)
		}

		var animeBackdrops []db.AnimeImage
		if len(animeBackdropIDs) > 0 {
			animeBackdrops = make([]db.AnimeImage, len(animeBackdropIDs))

			for i, b := range animeBackdropIDs {
				backdrop, err := server.gojo.GetAnimeImage(ctx, b)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime movie backdrop image", err)
				}
				animeBackdrops[i] = backdrop
			}
		}

		animeLogoIDs, err := server.gojo.ListAnimeMovieLogoImages(ctx, req.AnimeID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shv1.ApiError("cannot get anime movie logos images IDs", err)
		}

		var animeLogos []db.AnimeImage
		if len(animeLogoIDs) > 0 {
			animeLogos = make([]db.AnimeImage, len(animeLogoIDs))

			for i, l := range animeLogoIDs {
				logo, err := server.gojo.GetAnimeImage(ctx, l)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime movie logo image", err)
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
		animeTrailerIDs, err := server.gojo.ListAnimeMovieTrailers(ctx, req.AnimeID)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return shv1.ApiError("cannot get anime movie trailers IDs", err)
		}

		var animeTrailers []db.AnimeTrailer
		if len(animeTrailerIDs) > 0 {
			animeTrailers = make([]db.AnimeTrailer, len(animeTrailerIDs))

			for i, t := range animeTrailerIDs {
				trailer, err := server.gojo.GetAnimeTrailer(ctx, t.TrailerID)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime movie trailer", err)
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

func validateGetFullAnimeMovieRequest(req *ampbv1.GetFullAnimeMovieRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shv1.FieldViolation("animeID", err))
	}

	if err := utils.ValidateInt(int64(req.GetLanguageID())); err != nil {
		violations = append(violations, shv1.FieldViolation("languageID", err))
	}

	return violations
}
