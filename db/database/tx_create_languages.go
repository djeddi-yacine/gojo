package db

import (
	"context"
)

type CreateLanguagesTxParams struct {
	CreateLanguageParams []CreateLanguageParams
}

type CreateLanguagesTxResult struct {
	Languages []Language
}

func (gojo *SQLGojo) CreateLanguagesTx(ctx context.Context, arg CreateLanguagesTxParams) (CreateLanguagesTxResult, error) {
	var result CreateLanguagesTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		for _, p := range arg.CreateLanguageParams {
			l, err := q.CreateLanguage(ctx, p)
			if err != nil {
				return err
			}
			result.Languages = append(result.Languages, l)
		}

		return err
	})

	return result, err
}
