package repository

import (
	"context"

	"github.com/derangga/shopifyx/internal"
	"github.com/derangga/shopifyx/internal/entity"
	"github.com/derangga/shopifyx/internal/repository/query"
	"github.com/derangga/shopifyx/internal/repository/record"
	"github.com/jmoiron/sqlx"
)

type payment struct {
	db *sqlx.DB
}

func NewPaymentRepository(db *sqlx.DB) internal.PaymentRepository {
	return &payment{
		db: db,
	}
}

func (r *payment) Create(ctx context.Context, data *entity.Payment) error {
	_, err := handleTransaction(ctx, r.db, func(ctx context.Context, tx *sqlx.Tx) (*entity.Payment, error) {
		paymentRecord := record.PaymentEntityToRecord(data)

		stmt, err := tx.PrepareNamedContext(ctx, query.QueryInsertPayment)
		if err != nil {
			return nil, err
		}

		row := stmt.QueryRowContext(ctx, paymentRecord)
		if row.Err() != nil {
			return nil, row.Err()
		}

		return nil, nil
	})

	return err
}
