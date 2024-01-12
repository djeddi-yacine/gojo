package asapiv1

import (
	"context"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *AnimeSerieServer) GetAnimeSeasonEpisodes(ctx context.Context, req *aspbv1.GetAnimeSeasonEpisodesRequest) (*aspbv1.GetAnimeSeasonEpisodesResponse, error) {
	var err error

	_, err = shv1.AuthorizeUser(ctx, server.tokenMaker, utils.AllRolls)
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	violations := validateGetAnimeSeasonEpisodesRequest(req)
	if violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	cache := &ping.CacheKey{
		ID:     req.GetSeasonID(),
		Target: ping.AnimeSeason,
	}

	arg := db.ListAnimeSeasonEpisodesParams{
		SeasonID: req.GetSeasonID(),
		Limit:    req.GetPageSize(),
		Offset:   (req.GetPageNumber() - 1) * req.GetPageSize(),
	}

	var eIDs []int64
	if err = server.ping.Handle(ctx, cache.Episodes(arg.Limit, arg.Offset), &eIDs, func() error {
		eIDs, err = server.gojo.ListAnimeSeasonEpisodes(ctx, arg)
		if err != nil {
			return shv1.ApiError("failed to list all anime season episodes", err)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	res := &aspbv1.GetAnimeSeasonEpisodesResponse{}

	if len(eIDs) > 0 {
		res.SeasonEpisode = make([]*aspbv1.AnimeEpisodeResponse, len(eIDs))
		episodes := make([]db.AnimeSerieEpisode, len(eIDs))
		for i, v := range eIDs {
			cache = &ping.CacheKey{
				ID:     v,
				Target: ping.AnimeEpisode,
			}

			if err = server.ping.Handle(ctx, cache.Main(), &episodes[i], func() error {
				episodes[i], err = server.gojo.GetAnimeEpisodeByEpisodeID(ctx, v)
				if err != nil {
					return shv1.ApiError("failed to get anime season episodes", err)
				}

				return nil
			}); err != nil {
				return nil, err
			}

			res.SeasonEpisode[i] = convertAnimeEpisode(episodes[i])
		}
	}

	return res, nil
}

func validateGetAnimeSeasonEpisodesRequest(req *aspbv1.GetAnimeSeasonEpisodesRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetSeasonID()); err != nil {
		violations = append(violations, shv1.FieldViolation("seasonID", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageNumber())); err != nil {
		violations = append(violations, shv1.FieldViolation("pageNumber", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageSize())); err != nil {
		violations = append(violations, shv1.FieldViolation("pageSize", err))
	}

	return violations
}
