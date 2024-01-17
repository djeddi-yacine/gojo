package db

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/ping"
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
	CreateAnimeMovieTitlesTx(ctx context.Context, arg CreateAnimeMovieTitlesTxParams) (CreateAnimeMovieTitlesTxResult, error)
	QueryAnimeMovieTx(ctx context.Context, arg QueryAnimeMovieTxParams) (QueryAnimeMovieTxResult, error)
	CreateAnimeSeasonTitlesTx(ctx context.Context, arg CreateAnimeSeasonTitlesTxParams) (CreateAnimeSeasonTitlesTxResult, error)
	QueryAnimeSeasonTx(ctx context.Context, arg QueryAnimeSeasonTxParams) (QueryAnimeSeasonTxResult, error)
	CreateAnimeMovieTagTx(ctx context.Context, arg CreateAnimeMovieTagTxParams) (CreateAnimeMovieTagTxResult, error)
	CreateAnimeSeasonTagTx(ctx context.Context, arg CreateAnimeSeasonTagTxParams) (CreateAnimeSeasonTagTxResult, error)
	CreateAnimeEpisodeServerTx(ctx context.Context, episodeID int64) (AnimeEpisodeServer, error)
	LoginUserTx(ctx context.Context, arg LoginUserTxParams) (User, error)
	CreateActorsTx(ctx context.Context, arg []CreateActorParams) ([]Actor, error)
	CreateGenresTx(ctx context.Context, arg []string) ([]Genre, error)
	CreateStudiosTx(ctx context.Context, arg []string) ([]Studio, error)
	CreateLanguagesTx(ctx context.Context, arg []CreateLanguageParams) ([]Language, error)
	CreateAnimeMovieCharactersTx(ctx context.Context, arg CreateAnimeMovieCharactersTxParams) (CreateAnimeMovieCharactersTxResult, error)
	CreateAnimeSeasonCharactersTx(ctx context.Context, arg CreateAnimeSeasonCharactersTxParams) (CreateAnimeSeasonCharactersTxResult, error)
	ListGenresTx(ctx context.Context, arg []int32) ([]Genre, error)
	ListStudiosTx(ctx context.Context, arg []int32) ([]Studio, error)
	ListAnimeCharacetrsTx(ctx context.Context, arg []int64) ([]AnimeCharactersAndActors, error)
	ListAnimeTagsTx(ctx context.Context, arg []int64) ([]AnimeTag, error)
	ListAnimeImagesTx(ctx context.Context, arg []int64) ([]AnimeImage, error)
	ListAnimeTrailersTx(ctx context.Context, arg []int64) ([]AnimeTrailer, error)
	GetAllStudiosTx(ctx context.Context, arg ListStudiosParams) ([]Studio, error)
	GetAllLanguagesTx(ctx context.Context, arg ListLanguagesParams) ([]Language, error)
	GetAllGenresTx(ctx context.Context, arg ListGenresParams) ([]Genre, error)
	GetAllActorsTx(ctx context.Context, arg ListActorsParams) ([]Actor, error)
}

type SQLGojo struct {
	*Queries
	coonPool *pgxpool.Pool
	ping     *ping.PingSystem
}

func NewGojo(coonPool *pgxpool.Pool, ping *ping.PingSystem) Gojo {
	return &SQLGojo{
		coonPool: coonPool,
		Queries:  New(coonPool),
		ping:     ping,
	}
}
