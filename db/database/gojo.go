package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Gojo interface {
	Querier
	CreateAnimeMovieTx(ctx context.Context, arg CreateAnimeMovieTxParams) (CreateAnimeMovieTxResult, error)
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
