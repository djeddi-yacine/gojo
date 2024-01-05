package asapiv1

import (
	"context"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
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

	_, err = server.gojo.GetAnimeSeason(ctx, req.GetSeasonID())
	if err != nil {
		return nil, shv1.ApiError("failed to get the anime season", err)
	}

	arg := db.ListAnimeSeasonEpisodesParams{
		SeasonID: req.GetSeasonID(),
		Limit:    req.GetPageSize(),
		Offset:   (req.GetPageNumber() - 1) * req.GetPageSize(),
	}

	DBSeasonEpisodeIDs, err := server.gojo.ListAnimeSeasonEpisodes(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to list anime season episodes IDs", err)
	}

	DBSeasonEpisodes := make([]db.AnimeSerieEpisode, len(DBSeasonEpisodeIDs))
	for i, e := range DBSeasonEpisodeIDs {
		DBSeasonEpisodes[i], err = server.gojo.GetAnimeEpisodeByEpisodeID(ctx, e.EpisodeID)
		if err != nil {
			return nil, shv1.ApiError("failed to get anime season episodes", err)
		}
	}

	var PBSeasonEpisodes []*aspbv1.AnimeEpisodeResponse
	for _, e := range DBSeasonEpisodes {
		PBSeasonEpisodes = append(PBSeasonEpisodes, convertAnimeEpisode(e))
	}

	res := &aspbv1.GetAnimeSeasonEpisodesResponse{
		SeasonEpisode: PBSeasonEpisodes,
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
