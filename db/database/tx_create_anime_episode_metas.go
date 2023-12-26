package db

import (
	"context"
	"errors"
)

type CreateAnimeEpisodeMetasTxParams struct {
	EpisodeID    int64
	EpisodeMetas []AnimeMetaTxParam
}

type CreateAnimeEpisodeMetasTxResult struct {
	AnimeEpisodeMetas []AnimeMetaTxResult
}

func (gojo *SQLGojo) CreateAnimeEpisodeMetasTx(ctx context.Context, arg CreateAnimeEpisodeMetasTxParams) (CreateAnimeEpisodeMetasTxResult, error) {
	var result CreateAnimeEpisodeMetasTxResult
	var err error

	err = gojo.execTx(ctx, func(q *Queries) error {
		_, err = q.GetAnimeEpisodeByEpisodeID(ctx, arg.EpisodeID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if arg.EpisodeMetas != nil {
			result.AnimeEpisodeMetas = make([]AnimeMetaTxResult, len(arg.EpisodeMetas))

			for i, m := range arg.EpisodeMetas {
				lang, err := q.GetLanguage(ctx, m.LanguageID)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				meta, err := q.CreateMeta(ctx, m.CreateMetaParams)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				EpisodeMetaArg := CreateAnimeEpisodeMetaParams{
					EpisodeID:  arg.EpisodeID,
					LanguageID: lang.ID,
					MetaID:     meta.ID,
				}

				_, err = q.CreateAnimeEpisodeMeta(ctx, EpisodeMetaArg)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.AnimeEpisodeMetas[i] = AnimeMetaTxResult{
					Meta:       meta,
					LanguageID: m.LanguageID,
				}

			}
		} else {
			return errors.New("create one meta at least")
		}

		return err
	})

	return result, err
}
