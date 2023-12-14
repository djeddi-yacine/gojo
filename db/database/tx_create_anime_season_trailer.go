package db

import (
	"context"
)

type CreateAnimeSeasonTrailerTxParams struct {
	SeasonID             int64
	SeasonTrailersParams []CreateAnimeTrailerParams
}

type CreateAnimeSeasonTrailerTxResult struct {
	AnimeSeason    AnimeSerieSeason
	SeasonTrailers []AnimeTrailer
}

func (gojo *SQLGojo) CreateAnimeSeasonTrailerTx(ctx context.Context, arg CreateAnimeSeasonTrailerTxParams) (CreateAnimeSeasonTrailerTxResult, error) {
	var result CreateAnimeSeasonTrailerTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		season, err := q.GetAnimeSerieSeason(ctx, arg.SeasonID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeSeason = season

		if arg.SeasonTrailersParams != nil {
			if len(arg.SeasonTrailersParams) > 0 {
				var trailer CreateAnimeTrailerParams
				trailersArg := make([]CreateAnimeSeasonTrailerParams, len(arg.SeasonTrailersParams))
				result.SeasonTrailers = make([]AnimeTrailer, len(arg.SeasonTrailersParams))

				for i, t := range arg.SeasonTrailersParams {
					trailer = CreateAnimeTrailerParams{
						IsOfficial: t.IsOfficial,
						HostName:   t.HostName,
						HostKey:    t.HostKey,
					}

					tr, err := q.CreateAnimeTrailer(ctx, trailer)
					if err != nil {
						ErrorSQL(err)
						return err
					}

					result.SeasonTrailers[i] = tr
					trailersArg[i].SeasonID = season.ID
					trailersArg[i].TrailerID = tr.ID
				}

				for _, amt := range trailersArg {
					_, err = q.CreateAnimeSeasonTrailer(ctx, amt)
					if err != nil {
						ErrorSQL(err)
						return err
					}
				}
			}
		}

		return err
	})

	return result, err
}
