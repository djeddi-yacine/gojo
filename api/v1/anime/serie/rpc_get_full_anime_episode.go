package asapiv1

import (
	"context"

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

func (server *AnimeSerieServer) GetFullAnimeEpisode(ctx context.Context, req *aspbv1.GetFullAnimeEpisodeRequest) (*aspbv1.GetFullAnimeEpisodeResponse, error) {
	var err error

	_, err = shv1.AuthorizeUser(ctx, server.tokenMaker, utils.AllRolls)
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	violations := validateGetFullAnimeEpisodeRequest(req)
	if violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	cache := &ping.CacheKey{
		ID:      req.EpisodeID,
		Target:  ping.AnimeEpisode,
		Version: ping.V1,
	}

	res := &aspbv1.GetFullAnimeEpisodeResponse{}

	if err = server.ping.Handle(ctx, cache.Main(), &res.AnimeEpisode, func() error {
		animeEpisode, err := server.gojo.GetAnimeEpisodeByEpisodeID(ctx, req.GetEpisodeID())
		if err != nil {
			return shv1.ApiError("failed to get the anime episode", err)
		}

		res.AnimeEpisode = convertAnimeEpisode(animeEpisode)
		return nil
	}); err != nil {
		return nil, err
	}

	if err = server.ping.Handle(ctx, cache.Meta(uint32(req.LanguageID)), &res.EpisodeMeta, func() error {
		_, err := server.gojo.GetLanguage(ctx, req.GetLanguageID())
		if err != nil {
			return shv1.ApiError("failed to get the language", err)
		}

		animeMeta, err := server.gojo.GetAnimeEpisodeMeta(ctx, db.GetAnimeEpisodeMetaParams{
			EpisodeID:  req.GetEpisodeID(),
			LanguageID: req.GetLanguageID(),
		})
		if err != nil {
			return shv1.ApiError("no anime episode found with this language ID", err)
		}

		if animeMeta > 0 {
			meta, err := server.gojo.GetMeta(ctx, animeMeta)
			if err != nil {
				return shv1.ApiError("failed to get anime episode metadata", err)
			}

			res.EpisodeMeta = &nfpbv1.AnimeMetaResponse{
				LanguageID: req.GetLanguageID(),
				Meta:       shv1.ConvertMeta(meta),
				CreatedAt:  timestamppb.New(meta.CreatedAt),
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	if err = server.ping.Handle(ctx, cache.Server(), &res.ServerID, func() error {
		sv, err := server.gojo.GetAnimeEpisodeServer(ctx, req.GetEpisodeID())
		if err != nil {
			if db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot get anime episode server ID", err)
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
			ss, err := server.gojo.ListAnimeEpisodeServerSubVideos(ctx, res.ServerID)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot list anime episode server sub videos", err)
			}

			subV := make([]db.AnimeEpisodeVideo, len(ss))
			for i, ks := range ss {
				subV[i], err = server.gojo.GetAnimeEpisodeVideo(ctx, ks.VideoID)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime episode server sub videos", err)
				}
			}

			st, err := server.gojo.ListAnimeEpisodeServerSubTorrents(ctx, res.ServerID)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot list anime episode server sub torrents", err)
			}

			subT := make([]db.AnimeEpisodeTorrent, len(st))
			for i, kst := range st {
				subT[i], err = server.gojo.GetAnimeEpisodeTorrent(ctx, kst.ServerID)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime episode server sub torrents", err)
				}
			}

			res.Sub = &aspbv1.AnimeEpisodeSubDataResponse{
				Videos:   convertAnimeEpisodeVideos(subV),
				Torrents: convertAnimeEpisodeTorrents(subT),
			}

			return nil
		}); err != nil {
			return nil, err
		}

		if err = server.ping.Handle(ctx, cache.Dub(), &res.Dub, func() error {
			sd, err := server.gojo.ListAnimeEpisodeServerDubVideos(ctx, res.ServerID)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot list anime episode server dub videos", err)
			}

			dubV := make([]db.AnimeEpisodeVideo, len(sd))
			for i, kd := range sd {
				dubV[i], err = server.gojo.GetAnimeEpisodeVideo(ctx, kd.VideoID)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime episode server dub videos", err)
				}
			}

			dt, err := server.gojo.ListAnimeEpisodeServerDubTorrents(ctx, res.ServerID)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot list anime episode server dub torrents", err)
			}

			dubT := make([]db.AnimeEpisodeTorrent, len(dt))
			for i, kdt := range dt {
				dubT[i], err = server.gojo.GetAnimeEpisodeTorrent(ctx, kdt.ServerID)
				if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
					return shv1.ApiError("cannot get anime episode server dub torrents", err)
				}
			}

			res.Dub = &aspbv1.AnimeEpisodeDubDataResponse{
				Videos:   convertAnimeEpisodeVideos(dubV),
				Torrents: convertAnimeEpisodeTorrents(dubT),
			}
			return nil
		}); err != nil {
			return nil, err
		}
	}

	return res, nil
}

func validateGetFullAnimeEpisodeRequest(req *aspbv1.GetFullAnimeEpisodeRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetEpisodeID()); err != nil {
		violations = append(violations, shv1.FieldViolation("episodeID", err))
	}

	if err := utils.ValidateInt(int64(req.GetLanguageID())); err != nil {
		violations = append(violations, shv1.FieldViolation("languageID", err))
	}

	return violations
}
