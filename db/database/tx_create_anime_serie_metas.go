package db

import (
	"context"
)

type CreateAnimeSerieMetaTxParam struct {
	LanguageID int32
	CreateMetaParams
}

type CreateAnimeSerieMetaTxResult struct {
	Language Language
	Meta     Meta
}

type CreateAnimeSerieMetasTxParams struct {
	AnimeID                       int64
	CreateAnimeSerieMetasTxParams []CreateAnimeSerieMetaTxParam
}

type CreateAnimeSerieMetasTxResult struct {
	CreateAnimeSerieMetasTxResults []CreateAnimeSerieMetaTxResult
}

func (gojo *SQLGojo) CreateAnimeSerieMetasTx(ctx context.Context, arg CreateAnimeSerieMetasTxParams) (CreateAnimeSerieMetasTxResult, error) {
	var result CreateAnimeSerieMetasTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error
		result.CreateAnimeSerieMetasTxResults = make([]CreateAnimeSerieMetaTxResult, len(arg.CreateAnimeSerieMetasTxParams))

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

			result.CreateAnimeSerieMetasTxResults[i].Meta, err = q.GetMeta(ctx, animeMeta.MetaID)
			if err != nil {
				ErrorSQL(err)
				return err
			}

			result.CreateAnimeSerieMetasTxResults[i].Language, err = q.GetLanguage(ctx, animeMeta.LanguageID)
			if err != nil {
				ErrorSQL(err)
				return err
			}

		}

		return err
	})

	return result, err
}
