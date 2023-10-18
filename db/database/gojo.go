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
	CreateAnimeMovieMetaTx(ctx context.Context, arg CreateAnimeMovieMetaTxParams) (CreateAnimeMovieMetaTxResult, error)
	CreateGenresTx(ctx context.Context, arg CreateGenresTxParams) (CreateGenresTxResult, error)
	CreateStudiosTx(ctx context.Context, arg CreateStudiosTxParams) (CreateStudiosTxResult, error)
	CreateLanguagesTx(ctx context.Context, arg CreateLanguagesTxParams) (CreateLanguagesTxResult, error)
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
