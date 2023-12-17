package db

import (
	"context"
)

type CreateAnimeSeasonTitleTxParams struct {
	SeasonID            int64
	AnimeOfficialTitles []CreateAnimeSeasonOfficialTitleParams
	AnimeShortTitles    []CreateAnimeSeasonShortTitleParams
	AnimeOtherTitles    []CreateAnimeSeasonOtherTitleParams
}

type CreateAnimeSeasonTitleTxResult struct {
	AnimeSeason         AnimeSerieSeason
	AnimeOfficialTitles []AnimeSeasonOfficialTitle
	AnimeShortTitles    []AnimeSeasonShortTitle
	AnimeOtherTitles    []AnimeSeasonOtherTitle
}

func (gojo *SQLGojo) CreateAnimeSeasonTitleTx(ctx context.Context, arg CreateAnimeSeasonTitleTxParams) (CreateAnimeSeasonTitleTxResult, error) {
	var result CreateAnimeSeasonTitleTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		anime, err := q.GetAnimeSeason(ctx, arg.SeasonID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeSeason = anime

		if arg.AnimeOfficialTitles != nil {
			if len(arg.AnimeOfficialTitles) > 0 {
				var titleArg CreateAnimeSeasonOfficialTitleParams
				result.AnimeOfficialTitles = make([]AnimeSeasonOfficialTitle, len(arg.AnimeOfficialTitles))

				for i, t := range arg.AnimeOfficialTitles {
					titleArg = CreateAnimeSeasonOfficialTitleParams{
						SeasonID:  t.SeasonID,
						TitleText: t.TitleText,
					}

					title, err := q.CreateAnimeSeasonOfficialTitle(ctx, titleArg)
					if err != nil {
						ErrorSQL(err)
						return err
					}

					result.AnimeOfficialTitles[i] = title
				}
			}
		}

		if arg.AnimeShortTitles != nil {
			if len(arg.AnimeShortTitles) > 0 {
				var titleArg CreateAnimeSeasonShortTitleParams
				result.AnimeShortTitles = make([]AnimeSeasonShortTitle, len(arg.AnimeShortTitles))

				for i, t := range arg.AnimeShortTitles {
					titleArg = CreateAnimeSeasonShortTitleParams{
						SeasonID:  t.SeasonID,
						TitleText: t.TitleText,
					}

					title, err := q.CreateAnimeSeasonShortTitle(ctx, titleArg)
					if err != nil {
						ErrorSQL(err)
						return err
					}

					result.AnimeShortTitles[i] = title
				}
			}
		}

		if arg.AnimeOtherTitles != nil {
			if len(arg.AnimeOtherTitles) > 0 {
				var titleArg CreateAnimeSeasonOtherTitleParams
				result.AnimeOtherTitles = make([]AnimeSeasonOtherTitle, len(arg.AnimeOtherTitles))

				for i, t := range arg.AnimeOtherTitles {
					titleArg = CreateAnimeSeasonOtherTitleParams{
						SeasonID:  t.SeasonID,
						TitleText: t.TitleText,
					}

					title, err := q.CreateAnimeSeasonOtherTitle(ctx, titleArg)
					if err != nil {
						ErrorSQL(err)
						return err
					}

					result.AnimeOtherTitles[i] = title
				}
			}
		}

		return err
	})

	return result, err
}
