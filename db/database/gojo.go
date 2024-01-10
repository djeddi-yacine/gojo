package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Gojo interface {
	Querier
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error)
	CreateAnimeMovieInfoTx(ctx context.Context, arg CreateAnimeMovieInfoTxParams) (CreateAnimeMovieInfoTxResult, error)
	CreateAnimeSeasonInfoTx(ctx context.Context, arg CreateAnimeSeasonInfoTxParams) (CreateAnimeSeasonInfoTxResult, error)
	CreateAnimeMovieMetasTx(ctx context.Context, arg CreateAnimeMovieMetasTxParams) (CreateAnimeMovieMetasTxResult, error)
	CreateAnimeSerieMetasTx(ctx context.Context, arg CreateAnimeSerieMetasTxParams) (CreateAnimeSerieMetasTxResult, error)
	CreateAnimeMovieTx(ctx context.Context, arg CreateAnimeMovieTxParams) (CreateAnimeMovieTxResult, error)
	CreateAnimeSerieTx(ctx context.Context, arg CreateAnimeSerieTxParams) (CreateAnimeSerieTxResult, error)
	CreateAnimeMovieResourceTx(ctx context.Context, arg CreateAnimeMovieResourceTxParams) (CreateAnimeMovieResourceTxResult, error)
	CreateAnimeMovieLinkTx(ctx context.Context, arg CreateAnimeMovieLinkTxParams) (CreateAnimeMovieLinkTxResult, error)
	CreateAnimeSerieLinkTx(ctx context.Context, arg CreateAnimeSerieLinkTxParams) (CreateAnimeSerieLinkTxResult, error)
	CreateAnimeSeasonResourceTx(ctx context.Context, arg CreateAnimeSeasonResourceTxParams) (CreateAnimeSeasonResourceTxResult, error)
	CreateAnimeMovieDataTx(ctx context.Context, arg CreateAnimeMovieDataTxParams) (CreateAnimeMovieDataTxResult, error)
	CreateAnimeSeasonTx(ctx context.Context, arg CreateAnimeSeasonTxParams) (CreateAnimeSeasonTxResult, error)
	CreateAnimeSeasonMetasTx(ctx context.Context, arg CreateAnimeSeasonMetasTxParams) (CreateAnimeSeasonMetasTxResult, error)
	CreateAnimeEpisodeTx(ctx context.Context, arg CreateAnimeEpisodeTxParams) (CreateAnimeEpisodeTxResult, error)
	CreateAnimeEpisodeMetasTx(ctx context.Context, arg CreateAnimeEpisodeMetasTxParams) (CreateAnimeEpisodeMetasTxResult, error)
	CreateAnimeEpisodeDataTx(ctx context.Context, arg CreateAnimeEpisodeDataTxParams) (CreateAnimeEpisodeDataTxResult, error)
	CreateAnimeMovieImageTx(ctx context.Context, arg CreateAnimeMovieImageTxParams) (CreateAnimeMovieImageTxResult, error)
	CreateAnimeSerieImageTx(ctx context.Context, arg CreateAnimeSerieImageTxParams) (CreateAnimeSerieImageTxResult, error)
	CreateAnimeSeasonImageTx(ctx context.Context, arg CreateAnimeSeasonImageTxParams) (CreateAnimeSeasonImageTxResult, error)
	CreateAnimeMovieTrailerTx(ctx context.Context, arg CreateAnimeMovieTrailerTxParams) (CreateAnimeMovieTrailerTxResult, error)
	CreateAnimeSerieTrailerTx(ctx context.Context, arg CreateAnimeSerieTrailerTxParams) (CreateAnimeSerieTrailerTxResult, error)
	CreateAnimeSeasonTrailerTx(ctx context.Context, arg CreateAnimeSeasonTrailerTxParams) (CreateAnimeSeasonTrailerTxResult, error)
	CreateAnimeMovieTitleTx(ctx context.Context, arg CreateAnimeMovieTitleTxParams) (CreateAnimeMovieTitleTxResult, error)
	QueryAnimeMovieTx(ctx context.Context, arg QueryAnimeMovieTxParams) (QueryAnimeMovieTxResult, error)
	CreateAnimeSeasonTitleTx(ctx context.Context, arg CreateAnimeSeasonTitleTxParams) (CreateAnimeSeasonTitleTxResult, error)
	QueryAnimeSeasonTx(ctx context.Context, arg QueryAnimeSeasonTxParams) (QueryAnimeSeasonTxResult, error)
	CreateAnimeMovieTagTx(ctx context.Context, arg CreateAnimeMovieTagTxParams) (CreateAnimeMovieTagTxResult, error)
	CreateAnimeSeasonTagTx(ctx context.Context, arg CreateAnimeSeasonTagTxParams) (CreateAnimeSeasonTagTxResult, error)
	CreateAnimeEpisodeServerTx(ctx context.Context, episodeID int64) (AnimeEpisodeServer, error)
	LoginUserTx(ctx context.Context, arg LoginUserTxParams) (User, error)
	CreateActorsTx(ctx context.Context, arg []CreateActorParams) ([]Actor, error)
	CreateGenresTx(ctx context.Context, arg []string) ([]Genre, error)
	CreateStudiosTx(ctx context.Context, arg []string) ([]Studio, error)
	CreateLanguagesTx(ctx context.Context, arg []CreateLanguageParams) ([]Language, error)
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
