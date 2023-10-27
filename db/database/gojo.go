package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Gojo interface {
	Querier
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error)
	AddInfoAnimeMovieTx(ctx context.Context, arg AddInfoAnimeMovieTxParams) (AddInfoAnimeMovieTxResult, error)
	AddInfoAnimeSerieTx(ctx context.Context, arg AddInfoAnimeSerieTxParams) (AddInfoAnimeSerieTxResult, error)
	CreateAnimeMovieMetasTx(ctx context.Context, arg CreateAnimeMovieMetasTxParams) (CreateAnimeMovieMetasTxResult, error)
	CreateAnimeSerieMetasTx(ctx context.Context, arg CreateAnimeSerieMetasTxParams) (CreateAnimeSerieMetasTxResult, error)
	CreateGenresTx(ctx context.Context, arg CreateGenresTxParams) (CreateGenresTxResult, error)
	CreateStudiosTx(ctx context.Context, arg CreateStudiosTxParams) (CreateStudiosTxResult, error)
	CreateLanguagesTx(ctx context.Context, arg CreateLanguagesTxParams) (CreateLanguagesTxResult, error)
	RenewSessionTx(ctx context.Context, arg RenewSessionTxParams) (RenewSessionTxResult, error)
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
