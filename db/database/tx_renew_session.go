package db

import (
	"context"
)

type RenewSessionTxParams struct {
	CreateSessionParams
	AfterRenew func(username string) error
}

type RenewSessionTxResult struct {
	Session
}

func (gojo *SQLGojo) RenewSessionTx(ctx context.Context, arg RenewSessionTxParams) (RenewSessionTxResult, error) {
	var result RenewSessionTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error
		var isHere bool

		username, err := q.UpdateSession(ctx, UpdateSessionParams{
			Username:  arg.CreateSessionParams.Username,
			IsBlocked: true,
		})
		if err != nil {
			if err != ErrRecordNotFound {
				return err
			}
			err = nil
		} else {
			isHere = true
		}

		result.Session, err = q.CreateSession(ctx, arg.CreateSessionParams)
		if err != nil {
			return err
		}

		if isHere {
			return arg.AfterRenew(username)
		}

		return err
	})

	return result, err
}
