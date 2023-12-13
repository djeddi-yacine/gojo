package db

import (
	"context"
)

type CreateAnimeMovieLinkTxParams struct {
	AnimeID               int64
	CreateAnimeLinkParams CreateAnimeLinkParams
}

type CreateAnimeMovieLinkTxResult struct {
	AnimeMovie AnimeMovie
	AnimeLink  AnimeLink
}

func (gojo *SQLGojo) CreateAnimeMovieLinkTx(ctx context.Context, arg CreateAnimeMovieLinkTxParams) (CreateAnimeMovieLinkTxResult, error) {
	var result CreateAnimeMovieLinkTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		anime, err := q.GetAnimeMovie(ctx, arg.AnimeID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		link, err := q.CreateAnimeLink(ctx, arg.CreateAnimeLinkParams)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		arg := CreateAnimeMovieLinkParams{
			AnimeID: arg.AnimeID,
			LinkID:  link.ID,
		}

		_, err = q.CreateAnimeMovieLink(ctx, arg)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeLink = link
		result.AnimeMovie = anime

		return err
	})

	return result, err
}
