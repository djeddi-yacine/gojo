package db

import (
	"context"
)

type CreateAnimeSerieTxParams struct {
	CreateAnimeSerieParams CreateAnimeSerieParams
	CreateAnimeLinkParams  CreateAnimeLinkParams
}

type CreateAnimeSerieTxResult struct {
	AnimeSerie AnimeSerie
	AnimeLink  AnimeLink
}

func (gojo *SQLGojo) CreateAnimeSerieTx(ctx context.Context, arg CreateAnimeSerieTxParams) (CreateAnimeSerieTxResult, error) {
	var result CreateAnimeSerieTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		anime, err := q.CreateAnimeSerie(ctx, arg.CreateAnimeSerieParams)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		link, err := q.CreateAnimeLink(ctx, arg.CreateAnimeLinkParams)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		larg := CreateAnimeSerieLinkParams{
			AnimeID: anime.ID,
			LinkID:  link.ID,
		}

		_, err = q.CreateAnimeSerieLink(ctx, larg)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeSerie = anime
		result.AnimeLink = link

		return err
	})

	return result, err
}
