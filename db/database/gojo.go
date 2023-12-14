package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Gojo interface {
	Querier
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error)
	AddAnimeMovieInfoTx(ctx context.Context, arg AddAnimeMovieInfoTxParams) (AddAnimeMovieInfoTxResult, error)
	CreateAnimeSeasonInfoTx(ctx context.Context, arg CreateAnimeSeasonInfoTxParams) (CreateAnimeSeasonInfoTxResult, error)
	CreateAnimeMovieMetasTx(ctx context.Context, arg CreateAnimeMovieMetasTxParams) (CreateAnimeMovieMetasTxResult, error)
	CreateAnimeSerieMetasTx(ctx context.Context, arg CreateAnimeSerieMetasTxParams) (CreateAnimeSerieMetasTxResult, error)
	CreateGenresTx(ctx context.Context, arg CreateGenresTxParams) (CreateGenresTxResult, error)
	CreateStudiosTx(ctx context.Context, arg CreateStudiosTxParams) (CreateStudiosTxResult, error)
	CreateLanguagesTx(ctx context.Context, arg CreateLanguagesTxParams) (CreateLanguagesTxResult, error)
	RenewSessionTx(ctx context.Context, arg RenewSessionTxParams) (RenewSessionTxResult, error)
	CreateAnimeMovieTx(ctx context.Context, arg CreateAnimeMovieTxParams) (CreateAnimeMovieTxResult, error)
	CreateAnimeSerieTx(ctx context.Context, arg CreateAnimeSerieTxParams) (CreateAnimeSerieTxResult, error)
	CreateAnimeMovieResourceTx(ctx context.Context, arg CreateAnimeMovieResourceTxParams) (CreateAnimeMovieResourceTxResult, error)
	CreateAnimeMovieLinkTx(ctx context.Context, arg CreateAnimeMovieLinkTxParams) (CreateAnimeMovieLinkTxResult, error)
	CreateAnimeSerieLinkTx(ctx context.Context, arg CreateAnimeSerieLinkTxParams) (CreateAnimeSerieLinkTxResult, error)
	CreateAnimeSeasonResourceTx(ctx context.Context, arg CreateAnimeSeasonResourceTxParams) (CreateAnimeSeasonResourceTxResult, error)
	AddAnimeMovieDataTx(ctx context.Context, arg AddAnimeMovieDataTxParams) (AddAnimeMovieDataTxResult, error)
	CreateAnimeSeasonTx(ctx context.Context, arg CreateAnimeSeasonTxParams) (CreateAnimeSeasonTxResult, error)
	CreateAnimeSeasonMetasTx(ctx context.Context, arg CreateAnimeSeasonMetasTxParams) (CreateAnimeSeasonMetasTxResult, error)
	CreateAnimeEpisodeTx(ctx context.Context, arg CreateAnimeEpisodeTxParams) (CreateAnimeEpisodeTxResult, error)
	CreateAnimeEpisodeMetasTx(ctx context.Context, arg CreateAnimeEpisodeMetasTxParams) (CreateAnimeEpisodeMetasTxResult, error)
	CreateAnimeSerieDataTx(ctx context.Context, arg CreateAnimeSerieDataTxParams) (CreateAnimeSerieDataTxResult, error)
	CreateAnimeMovieImageTx(ctx context.Context, arg CreateAnimeMovieImageTxParams) (CreateAnimeMovieImageTxResult, error)
	CreateAnimeSerieImageTx(ctx context.Context, arg CreateAnimeSerieImageTxParams) (CreateAnimeSerieImageTxResult, error)
	CreateAnimeSeasonImageTx(ctx context.Context, arg CreateAnimeSeasonImageTxParams) (CreateAnimeSeasonImageTxResult, error)
	CreateAnimeMovieTrailerTx(ctx context.Context, arg CreateAnimeMovieTrailerTxParams) (CreateAnimeMovieTrailerTxResult, error)
	CreateAnimeSerieTrailerTx(ctx context.Context, arg CreateAnimeSerieTrailerTxParams) (CreateAnimeSerieTrailerTxResult, error)
	CreateAnimeSeasonTrailerTx(ctx context.Context, arg CreateAnimeSeasonTrailerTxParams) (CreateAnimeSeasonTrailerTxResult, error)
}

type SQLGojo struct {
	*Queries
	coonPool *pgxpool.Pool
}

func NewGojo(coonPool *pgxpool.Pool) Gojo {
	return &SQLGojo{
		coonPool: coonPool,
		Queries:  New(coonPool),
	}
}
