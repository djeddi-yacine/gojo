package db

import (
	"context"
)

type CreateAnimeSerieMetasTxParams struct {
	AnimeID                       int64
	CreateAnimeSerieMetasTxParams []AnimeMetaTxParam
}

type CreateAnimeSerieMetasTxResult struct {
	AnimeSerieMetasTxResults []AnimeMetaTxResult
}

func (gojo *SQLGojo) CreateAnimeSerieMetasTx(ctx context.Context, arg CreateAnimeSerieMetasTxParams) (CreateAnimeSerieMetasTxResult, error) {
	var result CreateAnimeSerieMetasTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error
		result.AnimeSerieMetasTxResults = make([]AnimeMetaTxResult, len(arg.CreateAnimeSerieMetasTxParams))

		for i, cs := range arg.CreateAnimeSerieMetasTxParams {
			meta, err := q.CreateMeta(ctx, cs.CreateMetaParams)
			if err != nil {
				ErrorSQL(err)
				return err
			}

			arg := CreateAnimeSerieMetaParams{
				AnimeID:    arg.AnimeID,
				LanguageID: cs.LanguageID,
				MetaID:     meta.ID,
			}
			animeMeta, err := q.CreateAnimeSerieMeta(ctx, arg)
			if err != nil {
				ErrorSQL(err)
				return err
			}

			result.AnimeSerieMetasTxResults[i].Meta, err = q.GetMeta(ctx, animeMeta.MetaID)
			if err != nil {
				ErrorSQL(err)
				return err
			}

			result.AnimeSerieMetasTxResults[i].LanguageID = animeMeta.LanguageID
		}

		return err
	})

	return result, err
}
