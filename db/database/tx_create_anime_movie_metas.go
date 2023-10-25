package db

import (
	"context"
)

type CreateAnimeMovieMetaTxParam struct {
	LanguageID int32
	CreateMetaParams
}

type CreateAnimeMovieMetaTxResult struct {
	Language Language
	Meta     Meta
}

type CreateAnimeMovieMetasTxParams struct {
	AnimeID                       int64
	CreateAnimeMovieMetasTxParams []CreateAnimeMovieMetaTxParam
}

type CreateAnimeMovieMetasTxResult struct {
	CreateAnimeMovieMetasTxResults []CreateAnimeMovieMetaTxResult
}

func (gojo *SQLGojo) CreateAnimeMovieMetasTx(ctx context.Context, arg CreateAnimeMovieMetasTxParams) (CreateAnimeMovieMetasTxResult, error) {
	var result CreateAnimeMovieMetasTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error
		result.CreateAnimeMovieMetasTxResults = make([]CreateAnimeMovieMetaTxResult, len(arg.CreateAnimeMovieMetasTxParams))

		for i, cm := range arg.CreateAnimeMovieMetasTxParams {
			meta, err := q.CreateMeta(ctx, cm.CreateMetaParams)
			if err != nil {
				ErrorSQL(err)
				return err
			}

			arg := CreateAnimeMovieMetaParams{
				AnimeID:    arg.AnimeID,
				LanguageID: cm.LanguageID,
				MetaID:     meta.ID,
			}
			animeMeta, err := q.CreateAnimeMovieMeta(ctx, arg)
			if err != nil {
				ErrorSQL(err)
				return err
			}

			result.CreateAnimeMovieMetasTxResults[i].Meta, err = q.GetMeta(ctx, animeMeta.MetaID)
			if err != nil {
				ErrorSQL(err)
				return err
			}

			result.CreateAnimeMovieMetasTxResults[i].Language, err = q.GetLanguage(ctx, animeMeta.LanguageID)
			if err != nil {
				ErrorSQL(err)
				return err
			}

		}

		return err
	})

	return result, err
}
