package db

import (
	"context"
)

type CreateAnimeMovieTitlesTxParams struct {
	AnimeID             int64
	AnimeOfficialTitles []CreateAnimeMovieOfficialTitleParams
	AnimeShortTitles    []CreateAnimeMovieShortTitleParams
	AnimeOtherTitles    []CreateAnimeMovieOtherTitleParams
}

type CreateAnimeMovieTitlesTxResult struct {
	AnimeOfficialTitles []AnimeMovieOfficialTitle
	AnimeShortTitles    []AnimeMovieShortTitle
	AnimeOtherTitles    []AnimeMovieOtherTitle
}

func (gojo *SQLGojo) CreateAnimeMovieTitlesTx(ctx context.Context, arg CreateAnimeMovieTitlesTxParams) (CreateAnimeMovieTitlesTxResult, error) {
	var result CreateAnimeMovieTitlesTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		if arg.AnimeOfficialTitles != nil {
			if len(arg.AnimeOfficialTitles) > 0 {
				var titleArg CreateAnimeMovieOfficialTitleParams
				result.AnimeOfficialTitles = make([]AnimeMovieOfficialTitle, len(arg.AnimeOfficialTitles))

				for i, t := range arg.AnimeOfficialTitles {
					titleArg = CreateAnimeMovieOfficialTitleParams{
						AnimeID:   t.AnimeID,
						TitleText: t.TitleText,
					}

					title, err := q.CreateAnimeMovieOfficialTitle(ctx, titleArg)
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
				var titleArg CreateAnimeMovieShortTitleParams
				result.AnimeShortTitles = make([]AnimeMovieShortTitle, len(arg.AnimeShortTitles))

				for i, t := range arg.AnimeShortTitles {
					titleArg = CreateAnimeMovieShortTitleParams{
						AnimeID:   t.AnimeID,
						TitleText: t.TitleText,
					}

					title, err := q.CreateAnimeMovieShortTitle(ctx, titleArg)
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
				var titleArg CreateAnimeMovieOtherTitleParams
				result.AnimeOtherTitles = make([]AnimeMovieOtherTitle, len(arg.AnimeOtherTitles))

				for i, t := range arg.AnimeOtherTitles {
					titleArg = CreateAnimeMovieOtherTitleParams{
						AnimeID:   t.AnimeID,
						TitleText: t.TitleText,
					}

					title, err := q.CreateAnimeMovieOtherTitle(ctx, titleArg)
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
