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
	AnimeEpisode      AnimeSerieEpisode
	AnimeEpisodeMetas []AnimeMetaTxResult
}

func (gojo *SQLGojo) CreateAnimeEpisodeMetasTx(ctx context.Context, arg CreateAnimeEpisodeMetasTxParams) (CreateAnimeEpisodeMetasTxResult, error) {
	var result CreateAnimeEpisodeMetasTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		Episode, err := q.GetAnimeEpisodeByEpisodeID(ctx, arg.EpisodeID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeEpisode = Episode

		if arg.EpisodeMetas != nil {
			var metaArg CreateMetaParams
			var EpisodeMetaArg CreateAnimeEpisodeMetaParams
			result.AnimeEpisodeMetas = make([]AnimeMetaTxResult, len(arg.EpisodeMetas))

			for i, m := range arg.EpisodeMetas {
				metaArg = CreateMetaParams{
					Title:    m.Title,
					Overview: m.Overview,
				}

				meta, err := q.CreateMeta(ctx, metaArg)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				EpisodeMetaArg = CreateAnimeEpisodeMetaParams{
					EpisodeID:  Episode.ID,
					LanguageID: m.LanguageID,
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
