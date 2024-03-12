package repository

import (
	"context"

	"github.com/derangga/shopifyx/internal"
	pkgcontext "github.com/derangga/shopifyx/internal/pkg/context"
	"github.com/jmoiron/sqlx"
)

type UnitOfWork struct {
	DB *sqlx.DB
}

func NewUnitOfWork(db *sqlx.DB) internal.UnitOfWork {
	return &UnitOfWork{
		DB: db,
	}
}

func (r *UnitOfWork) BeginContext(ctx context.Context) (context.Context, *sqlx.Tx, error) {
	ctx, tx, _, err := pkgcontext.GetOrCreateTransactionContext(ctx, r.DB)
	return ctx, tx, err
}

func (r *UnitOfWork) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	outCtx, tx, txCreated, err := pkgcontext.GetOrCreateTransactionContext(ctx, r.DB)
	if err != nil {
		return err
	}

	err = fn(outCtx)

	if !txCreated {
		return err
	}

	defer tx.Rollback()
	if err != nil {
		return err
	}

	return tx.Commit()
}

func handleTransaction[T any](ctx context.Context, db *sqlx.DB, fn func(ctx context.Context, tx *sqlx.Tx) (*T, error)) (*T, error) {
	_, tx, newTx, err := pkgcontext.GetOrCreateTransactionContext(ctx, db)
	if err != nil {
		return nil, err
	}

	if newTx {
		defer tx.Rollback()
	}

	res, err := fn(ctx, tx)
	if err != nil {
		return nil, err
	}

	// If the transaction was already there, we don't need to commit it
	if !newTx {
		return res, nil
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return res, nil
}
