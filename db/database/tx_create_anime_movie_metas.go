package db

import (
	"context"
)

type CreateAnimeMovieMetasTxParams struct {
	AnimeID                       int64
	CreateAnimeMovieMetasTxParams []AnimeMetaTxParam
}

type CreateAnimeMovieMetasTxResult struct {
	AnimeMovieMetasTxResults []AnimeMetaTxResult
}

func (gojo *SQLGojo) CreateAnimeMovieMetasTx(ctx context.Context, arg CreateAnimeMovieMetasTxParams) (CreateAnimeMovieMetasTxResult, error) {
	var result CreateAnimeMovieMetasTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error
		result.AnimeMovieMetasTxResults = make([]AnimeMetaTxResult, len(arg.CreateAnimeMovieMetasTxParams))

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

			result.AnimeMovieMetasTxResults[i].Meta, err = q.GetMeta(ctx, animeMeta.MetaID)
			if err != nil {
				ErrorSQL(err)
				return err
			}

			result.AnimeMovieMetasTxResults[i].LanguageID = animeMeta.LanguageID

			_, err = q.CreateAnimeMovieTranslationTitle(ctx, CreateAnimeMovieTranslationTitleParams{
				AnimeID:   arg.AnimeID,
				TitleText: meta.Title,
			})
			if err != nil {
				ErrorSQL(err)
				return err
			}
		}

		return err
	})

	return result, err
}
