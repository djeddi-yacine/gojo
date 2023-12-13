package db

import (
	"context"
)

type CreateAnimeSerieLinkTxParams struct {
	AnimeID               int64
	CreateAnimeLinkParams CreateAnimeLinkParams
}

type CreateAnimeSerieLinkTxResult struct {
	AnimeSerie AnimeSerie
	AnimeLink  AnimeLink
}

func (gojo *SQLGojo) CreateAnimeSerieLinkTx(ctx context.Context, arg CreateAnimeSerieLinkTxParams) (CreateAnimeSerieLinkTxResult, error) {
	var result CreateAnimeSerieLinkTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		anime, err := q.GetAnimeSerie(ctx, arg.AnimeID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		link, err := q.CreateAnimeLink(ctx, arg.CreateAnimeLinkParams)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		arg := CreateAnimeSerieLinkParams{
			AnimeID: arg.AnimeID,
			LinkID:  link.ID,
		}

		_, err = q.CreateAnimeSerieLink(ctx, arg)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeLink = link
		result.AnimeSerie = anime

		return err
	})

	return result, err
}
