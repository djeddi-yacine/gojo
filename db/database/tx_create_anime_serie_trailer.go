package db

import (
	"context"
)

type CreateAnimeSerieTrailerTxParams struct {
	AnimeID             int64
	AnimeTrailersParams []CreateAnimeTrailerParams
}

type CreateAnimeSerieTrailerTxResult struct {
	AnimeSerie    AnimeSerie
	AnimeTrailers []AnimeTrailer
}

func (gojo *SQLGojo) CreateAnimeSerieTrailerTx(ctx context.Context, arg CreateAnimeSerieTrailerTxParams) (CreateAnimeSerieTrailerTxResult, error) {
	var result CreateAnimeSerieTrailerTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		anime, err := q.GetAnimeSerie(ctx, arg.AnimeID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeSerie = anime

		if arg.AnimeTrailersParams != nil {
			if len(arg.AnimeTrailersParams) > 0 {
				var trailer CreateAnimeTrailerParams
				trailersArg := make([]CreateAnimeSerieTrailerParams, len(arg.AnimeTrailersParams))
				result.AnimeTrailers = make([]AnimeTrailer, len(arg.AnimeTrailersParams))

				for i, t := range arg.AnimeTrailersParams {
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

					result.AnimeTrailers[i] = tr
					trailersArg[i].AnimeID = anime.ID
					trailersArg[i].TrailerID = tr.ID
				}

				for _, amt := range trailersArg {
					_, err = q.CreateAnimeSerieTrailer(ctx, amt)
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
