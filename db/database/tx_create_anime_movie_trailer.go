package db

import (
	"context"
)

type CreateAnimeMovieTrailerTxParams struct {
	AnimeID             int64
	AnimeTrailersParams []CreateAnimeTrailerParams
}

type CreateAnimeMovieTrailerTxResult struct {
	AnimeMovie    AnimeMovie
	AnimeTrailers []AnimeTrailer
}

func (gojo *SQLGojo) CreateAnimeMovieTrailerTx(ctx context.Context, arg CreateAnimeMovieTrailerTxParams) (CreateAnimeMovieTrailerTxResult, error) {
	var result CreateAnimeMovieTrailerTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		anime, err := q.GetAnimeMovie(ctx, arg.AnimeID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeMovie = anime

		if arg.AnimeTrailersParams != nil {
			if len(arg.AnimeTrailersParams) > 0 {
				var trailer CreateAnimeTrailerParams
				trailersArg := make([]CreateAnimeMovieTrailerParams, len(arg.AnimeTrailersParams))
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
					_, err = q.CreateAnimeMovieTrailer(ctx, amt)
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
