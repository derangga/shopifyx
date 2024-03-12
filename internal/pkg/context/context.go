package context

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type contextKey string

const transactionContextKey contextKey = "TransactionContextKey"

// GetOrCreateTransactionContext sets transaction to context
func GetOrCreateTransactionContext(
	ctx context.Context,
	db *sqlx.DB,
) (outCtx context.Context, tx *sqlx.Tx, created bool, err error) {
	tx, ok := ctx.Value(transactionContextKey).(*sqlx.Tx)
	if ok {
		return ctx, tx, false, nil
	}

	tx, err = db.BeginTxx(ctx, nil)
	if err != nil {
		return ctx, nil, false, err
	}

	outCtx = context.WithValue(ctx, transactionContextKey, tx)
	return outCtx, tx, true, err
}
