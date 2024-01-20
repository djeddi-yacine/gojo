package amapiv1

import (
	"context"

	av1 "github.com/dj-yacine-flutter/gojo/api/v1/anime"
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

func (server *AnimeMovieServer) GetOptionalFullAnimeMovie(ctx context.Context, req *ampbv1.GetOptionalFullAnimeMovieRequest) (*ampbv1.GetOptionalFullAnimeMovieResponse, error) {
	var err error

	_, err = shv1.AuthorizeUser(ctx, server.tokenMaker, utils.AllRolls)
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	violations := validateGetOptionalFullAnimeMovieRequest(req)
	if violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	cache := &ping.CacheKey{
		ID:     req.GetAnimeID(),
		Target: ping.AnimeMovie,
	}

	res := &ampbv1.GetOptionalFullAnimeMovieResponse{}

	var movie db.AnimeMovie
	if err = server.ping.Handle(ctx, cache.Main(), &movie, func() error {
		movie, err = server.gojo.GetAnimeMovie(ctx, req.GetAnimeID())
		if err != nil {
			return shv1.ApiError("cannot get anime movie", err)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	res.AnimeMovie = convertAnimeMovie(movie)

	var meta db.Meta
	if err = server.ping.Handle(ctx, cache.Meta(), &meta, func() error {
		animeMeta, err := server.gojo.GetAnimeMovieMeta(ctx, db.GetAnimeMovieMetaParams{
			AnimeID:    req.GetAnimeID(),
			LanguageID: req.GetLanguageID(),
		})
		if err != nil {
			return shv1.ApiError("no anime movie found with this language ID", err)
		}

		if animeMeta > 0 {
			meta, err = server.gojo.GetMeta(ctx, animeMeta)
			if err != nil {
				return shv1.ApiError("cannot get anime movie metadata", err)
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

	if req.GetWithResources() {
		var resources db.AnimeResource
		if err = server.ping.Handle(ctx, cache.Resources(), &resources, func() error {
			ID, err := server.gojo.GetAnimeMovieResource(ctx, req.GetAnimeID())
			if err != nil {
				if db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime movie resources", err)
				} else {
					return nil
				}
			}

			resources, err = server.gojo.GetAnimeResource(ctx, ID.ResourceID)
			if err != nil {
				if db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get resources data", err)
				} else {
					return nil
				}
			}

			return nil
		}); err != nil {
			return nil, err
		}

		res.AnimeResources = av1.ConvertAnimeResource(resources)
	}

	if req.GetWithGenres() {
		var gIDs []int32
		if err = server.ping.Handle(ctx, cache.Genre(), &gIDs, func() error {
			gIDs, err = server.gojo.ListAnimeMovieGenres(ctx, req.GetAnimeID())
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot get anime movie genres", err)
			}

			return nil
		}); err != nil {
			return nil, err
		}

		genres, err := server.gojo.ListGenresTx(ctx, gIDs)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return nil, shv1.ApiError("cannot list anime movie genres", err)
		}

		res.AnimeGenres = shv1.ConvertGenres(genres)
	}

	if req.GetWithStudios() {
		var sIDs []int32
		if err = server.ping.Handle(ctx, cache.Studio(), &sIDs, func() error {
			sIDs, err = server.gojo.ListAnimeMovieStudios(ctx, req.GetAnimeID())
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot get anime movie studios", err)
			}

			return nil
		}); err != nil {
			return nil, err
		}

		studios, err := server.gojo.ListStudiosTx(ctx, sIDs)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return nil, shv1.ApiError("cannot list anime movie studios", err)
		}

		res.AnimeStudios = shv1.ConvertStudios(studios)
	}

	var serverID int64
	if req.GetWithServer() {
		if err = server.ping.Handle(ctx, cache.Server(), &serverID, func() error {
			serverID, err = server.gojo.GetAnimeMovieServerByAnimeID(ctx, req.GetAnimeID())
			if err != nil {
				if db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime movie server ID", err)
				} else {
					return nil
				}
			}

			return nil
		}); err != nil {
			return nil, err
		}

		res.ServerID = &serverID
	}

	if serverID != 0 {
		if req.GetWithSub() {
			var subVideos []db.AnimeMovieVideo
			if err = server.ping.Handle(ctx, cache.SubV(), &subVideos, func() error {
				v, err := server.gojo.ListAnimeMovieServerSubVideos(ctx, serverID)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot list anime movie server sub videos", err)
				}

				if len(v) > 0 {
					subVideos = make([]db.AnimeMovieVideo, len(v))
					for i, x := range v {
						subVideos[i], err = server.gojo.GetAnimeMovieVideo(ctx, x.VideoID)
						if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
							return shv1.ApiError("cannot get anime movie server sub videos", err)
						}
					}
				}

				return nil
			}); err != nil {
				return nil, err
			}

			var subTorrents []db.AnimeMovieTorrent
			if err = server.ping.Handle(ctx, cache.SubT(), &subTorrents, func() error {
				v, err := server.gojo.ListAnimeMovieServerSubTorrents(ctx, serverID)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot list anime movie server sub torrents", err)
				}

				if len(v) > 0 {
					subTorrents = make([]db.AnimeMovieTorrent, len(v))
					for i, x := range v {
						subTorrents[i], err = server.gojo.GetAnimeMovieTorrent(ctx, x.TorrentID)
						if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
							return shv1.ApiError("cannot get anime movie server sub torrents", err)
						}
					}
				}

				return nil
			}); err != nil {
				return nil, err
			}

			res.Sub = &ashpbv1.AnimeSubDataResponse{
				Videos:   convertAnimeMovieVideos(subVideos),
				Torrents: convertAnimeMovieTorrents(subTorrents),
			}
		}

		if req.GetWithDub() {
			var dubVideos []db.AnimeMovieVideo
			if err = server.ping.Handle(ctx, cache.DubV(), &dubVideos, func() error {
				v, err := server.gojo.ListAnimeMovieServerDubVideos(ctx, serverID)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot list anime movie server dub videos", err)
				}

				if len(v) > 0 {
					dubVideos = make([]db.AnimeMovieVideo, len(v))
					for i, x := range v {
						dubVideos[i], err = server.gojo.GetAnimeMovieVideo(ctx, x.VideoID)
						if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
							return shv1.ApiError("cannot get anime movie server dub videos", err)
						}
					}
				}

				return nil
			}); err != nil {
				return nil, err
			}

			var dubTorrents []db.AnimeMovieTorrent
			if err = server.ping.Handle(ctx, cache.DubT(), &dubTorrents, func() error {
				v, err := server.gojo.ListAnimeMovieServerDubTorrents(ctx, serverID)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot list anime movie server dub torrents", err)
				}

				if len(v) > 0 {
					dubTorrents = make([]db.AnimeMovieTorrent, len(v))
					for i, x := range v {
						dubTorrents[i], err = server.gojo.GetAnimeMovieTorrent(ctx, x.TorrentID)
						if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
							return shv1.ApiError("cannot get anime movie server dub torrents", err)
						}
					}
				}

				return nil
			}); err != nil {
				return nil, err
			}

			res.Dub = &ashpbv1.AnimeDubDataResponse{
				Videos:   convertAnimeMovieVideos(dubVideos),
				Torrents: convertAnimeMovieTorrents(dubTorrents),
			}
		}
	}

	if req.GetWithLinks() {
		var link db.AnimeLink
		if err = server.ping.Handle(ctx, cache.Links(), &link, func() error {
			ID, err := server.gojo.GetAnimeMovieLink(ctx, req.GetAnimeID())
			if err != nil {
				if db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime movie links ID", err)
				} else {
					return nil
				}
			}

			link, err = server.gojo.GetAnimeLink(ctx, ID.ID)
			if err != nil {
				if db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime movie links", err)
				} else {
					return nil
				}
			}

			return nil
		}); err != nil {
			return nil, err
		}

		res.AnimeLinks = av1.ConvertAnimeLink(link)
	}

	if req.GetWithTags() {
		var tIDs []int64
		if err = server.ping.Handle(ctx, cache.Tags(), &tIDs, func() error {
			tIDs, err = server.gojo.ListAnimeMovieTags(ctx, req.GetAnimeID())
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot get anime movie tags IDs", err)
			}

			return nil
		}); err != nil {
			return nil, err
		}

		tags, err := server.gojo.ListAnimeTagsTx(ctx, tIDs)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return nil, shv1.ApiError("cannot get anime movie tag", err)
		}

		res.AnimeTags = av1.ConvertAnimeTags(tags)
	}

	if req.GetWithImages() {
		var pIDs []int64
		if err = server.ping.Handle(ctx, cache.Posters(), &pIDs, func() error {
			pIDs, err = server.gojo.ListAnimeMoviePosterImages(ctx, req.GetAnimeID())
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
			bIDs, err = server.gojo.ListAnimeMovieBackdropImages(ctx, req.GetAnimeID())
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
			lIDs, err = server.gojo.ListAnimeMovieLogoImages(ctx, req.GetAnimeID())
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

		res.AnimeImages = &ashpbv1.AnimeImageResponse{
			Posters:   av1.ConvertAnimeImages(posters),
			Backdrops: av1.ConvertAnimeImages(backdrops),
			Logos:     av1.ConvertAnimeImages(logos),
		}
	}

	if req.GetWithTrailer() {
		var rIDs []int64
		if err = server.ping.Handle(ctx, cache.Trailers(), &rIDs, func() error {
			rIDs, err = server.gojo.ListAnimeMovieTrailers(ctx, req.GetAnimeID())
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot get anime movie trailers IDs", err)
			}

			return nil
		}); err != nil {
			return nil, err
		}

		trailers, err := server.gojo.ListAnimeTrailersTx(ctx, rIDs)
		if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
			return nil, shv1.ApiError("cannot get anime movie trailers", err)
		}

		res.AnimeTrailers = av1.ConvertAnimeTrailers(trailers)
	}

	return res, nil
}

func validateGetOptionalFullAnimeMovieRequest(req *ampbv1.GetOptionalFullAnimeMovieRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shv1.FieldViolation("animeID", err))
	}

	if err := utils.ValidateInt(int64(req.GetLanguageID())); err != nil {
		violations = append(violations, shv1.FieldViolation("languageID", err))
	}

	return violations
}
