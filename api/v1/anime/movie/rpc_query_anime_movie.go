package amapiv1

import (
	"context"
	"fmt"
	"strconv"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/meilisearch/meilisearch-go"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *AnimeMovieServer) QueryAnimeMovie(ctx context.Context, req *ampbv1.QueryAnimeMovieRequest) (*ampbv1.QueryAnimeMovieResponse, error) {
	var err error

	_, err = shv1.AuthorizeUser(ctx, server.tokenMaker, utils.AllRolls)
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	violations := validateQueryAnimeMovieRequest(req)
	if violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	res := &ampbv1.QueryAnimeMovieResponse{}

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
		arg := db.QueryAnimeMovieTxParams{
			Query:  req.GetQuery(),
			Limit:  req.GetPageSize(),
			Offset: (req.GetPageNumber() - 1) * req.GetPageSize(),
		}

		data, err := server.gojo.QueryAnimeMovieTx(ctx, arg)
		if err != nil {
			return nil, shv1.ApiError("failed to query anime movies", err)
		}

		res.AnimeMovies = make([]*ampbv1.AnimeMovieResponse, len(data.AnimeMovies))
		for i, v := range data.AnimeMovies {
			res.AnimeMovies[i] = convertAnimeMovie(v)
		}

		return res, nil
	}

	res.AnimeMovies = make([]*ampbv1.AnimeMovieResponse, len(result.Hits))

	cache := ping.CacheKey{
		ID:     0,
		Target: ping.AnimeMovie,
	}

	var anime db.AnimeMovie

	for i, v := range result.Hits {
		id, err := strconv.Atoi(fmt.Sprint(v.(map[string]interface{})["ID"]))
		if err != nil {
			continue
		}

		cache.ID = int64(id)

		if err = server.ping.Handle(ctx, cache.Main(), &anime, func() error {
			anime, err = server.gojo.GetAnimeMovie(ctx, int64(id))
			if err != nil {
				return err
			}

			return nil
		}); err != nil {
			continue
		}

		res.AnimeMovies[i] = convertAnimeMovie(anime)
	}

	return res, nil
}

func validateQueryAnimeMovieRequest(req *ampbv1.QueryAnimeMovieRequest) (violations []*errdetails.BadRequest_FieldViolation) {
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
