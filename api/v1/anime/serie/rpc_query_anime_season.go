package asapiv1

import (
	"context"
	"fmt"
	"strconv"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/meilisearch/meilisearch-go"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *AnimeSerieServer) QueryAnimeSeason(ctx context.Context, req *aspbv1.QueryAnimeSeasonRequest) (*aspbv1.QueryAnimeSeasonResponse, error) {
	var err error

	_, err = shv1.AuthorizeUser(ctx, server.tokenMaker, utils.AllRolls)
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	violations := validateQueryAnimeSeasonRequest(req)
	if violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	res := &aspbv1.QueryAnimeSeasonResponse{}

	result, err := server.meilisearch.Search(req.Query,
		&meilisearch.SearchRequest{
			Page:                 int64(req.GetPageNumber()),
			HitsPerPage:          int64(req.GetPageSize()),
			AttributesToRetrieve: []string{"ID"},
			Filter:               fmt.Sprintf("ID != %s", req.Query),
			ShowMatchesPosition:  false,
			ShowRankingScore:     false,
			PlaceholderSearch:    false,
		})

	if err != nil || len(result.Hits) <= 0 {
		arg := db.QueryAnimeSeasonTxParams{
			Query:  req.GetQuery(),
			Limit:  req.GetPageSize(),
			Offset: (req.GetPageNumber() - 1) * req.GetPageSize(),
		}

		data, err := server.gojo.QueryAnimeSeasonTx(ctx, arg)
		if err != nil {
			return nil, shv1.ApiError("failed to query anime seasons", err)
		}

		res.AnimeSeasons = make([]*aspbv1.AnimeSeasonResponse, len(data.AnimeSeasons))
		for i, v := range data.AnimeSeasons {
			res.AnimeSeasons[i] = server.convertAnimeSeason(v)
		}

		return res, nil
	}

	res.AnimeSeasons = make([]*aspbv1.AnimeSeasonResponse, len(result.Hits))

	cache := ping.CacheKey{
		ID:     0,
		Target: ping.AnimeSeason,
	}

	var anime db.AnimeSeason

	for i, v := range result.Hits {
		id, err := strconv.Atoi(fmt.Sprint(v.(map[string]interface{})["ID"]))
		if err != nil {
			continue
		}

		cache.ID = int64(id)

		if err = server.ping.Handle(ctx, cache.Main(), &anime, func() error {
			anime, err = server.gojo.GetAnimeSeason(ctx, int64(id))
			if err != nil {
				return err
			}

			return nil
		}); err != nil {
			continue
		}

		res.AnimeSeasons[i] = server.convertAnimeSeason(anime)
	}

	return res, nil
}

func validateQueryAnimeSeasonRequest(req *aspbv1.QueryAnimeSeasonRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateString(req.GetQuery(), 1, 150); err != nil {
		violations = append(violations, shv1.FieldViolation("query", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageNumber())); err != nil {
		violations = append(violations, shv1.FieldViolation("pageNumber", err))
	}

	if err := utils.ValidateInt(int64(req.GetPageSize())); err != nil {
		violations = append(violations, shv1.FieldViolation("pageSize", err))
	}

	return violations
}
