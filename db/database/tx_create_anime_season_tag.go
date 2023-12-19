package db

import (
	"context"
)

type CreateAnimeSeasonTagTxParams struct {
	SeasonID   int64
	SeasonTags []string
}

type CreateAnimeSeasonTagTxResult struct {
	AnimeSeason AnimeSerieSeason
	SeasonTags  []AnimeTag
}

func (gojo *SQLGojo) CreateAnimeSeasonTagTx(ctx context.Context, arg CreateAnimeSeasonTagTxParams) (CreateAnimeSeasonTagTxResult, error) {
	var result CreateAnimeSeasonTagTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		season, err := q.GetAnimeSeason(ctx, arg.SeasonID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeSeason = season

		if arg.SeasonTags != nil {
			if len(arg.SeasonTags) > 0 {
				var tag AnimeTag
				tagsArg := make([]CreateAnimeSeasonTagParams, len(arg.SeasonTags))
				result.SeasonTags = make([]AnimeTag, len(arg.SeasonTags))

				for i, t := range arg.SeasonTags {
					tag, err = q.CreateAnimeTag(ctx, t)
					if err != nil {
						ErrorSQL(err)
						return err
					}

					result.SeasonTags[i] = tag
					tagsArg[i].SeasonID = season.ID
					tagsArg[i].TagID = tag.ID
				}

				for _, amt := range tagsArg {
					_, err = q.CreateAnimeSeasonTag(ctx, amt)
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
