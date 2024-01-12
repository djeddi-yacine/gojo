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

func (server *AnimeSerieServer) GetAnimeSerieSeasons(ctx context.Context, req *aspbv1.GetAnimeSerieSeasonsRequest) (*aspbv1.GetAnimeSerieSeasonsResponse, error) {
	var err error

	_, err = shv1.AuthorizeUser(ctx, server.tokenMaker, utils.AllRolls)
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	violations := validateGetAnimeSerieSeasonsRequest(req)
	if violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	_, err = server.gojo.GetAnimeSerie(ctx, req.GetAnimeID())
	if err != nil {
		return nil, shv1.ApiError("failed to get the anime serie", err)
	}

	cache := &ping.CacheKey{
		ID:     req.GetAnimeID(),
		Target: ping.AnimeSerie,
	}

	arg := db.ListAnimeSeasonsByAnimeIDParams{
		AnimeID: req.GetAnimeID(),
		Limit:   req.GetPageSize(),
		Offset:  (req.GetPageNumber() - 1) * req.GetPageSize(),
	}

	var sIDs []int64
	if err = server.ping.Handle(ctx, cache.Seasons(arg.Limit, arg.Offset), &sIDs, func() error {
		sIDs, err = server.gojo.ListAnimeSeasonsByAnimeID(ctx, arg)
		if err != nil {
			return shv1.ApiError("failed to list all anime serie seasons", err)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	res := &aspbv1.GetAnimeSerieSeasonsResponse{}

	if len(sIDs) > 0 {
		res.AnimeSeasons = make([]*aspbv1.AnimeSeasonResponse, len(sIDs))
		seasons := make([]db.AnimeSerieSeason, len(sIDs))
		for i, v := range sIDs {
			cache = &ping.CacheKey{
				ID:     v,
				Target: ping.AnimeSeason,
			}

			if err = server.ping.Handle(ctx, cache.Main(), &seasons[i], func() error {
				seasons[i], err = server.gojo.GetAnimeSeason(ctx, v)
				if err != nil {
					return shv1.ApiError("failed to get anime serie seasons", err)
				}

				return nil
			}); err != nil {
				return nil, err
			}

			res.AnimeSeasons[i] = convertAnimeSeason(seasons[i])
		}
	}

	return res, nil
}

func validateGetAnimeSerieSeasonsRequest(req *aspbv1.GetAnimeSerieSeasonsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shv1.FieldViolation("serieID", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageNumber())); err != nil {
		violations = append(violations, shv1.FieldViolation("pageNumber", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageSize())); err != nil {
		violations = append(violations, shv1.FieldViolation("pageSize", err))
	}

	return violations
}
