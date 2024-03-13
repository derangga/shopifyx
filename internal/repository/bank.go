package repository

import (
	"context"

	"github.com/derangga/shopifyx/internal"
	"github.com/derangga/shopifyx/internal/entity"
	"github.com/derangga/shopifyx/internal/repository/query"
	"github.com/derangga/shopifyx/internal/repository/record"
	"github.com/jmoiron/sqlx"
)

type bank struct {
	db *sqlx.DB
}

func NewBankRepository(db *sqlx.DB) internal.BankRepository {
	return &bank{
		db: db,
	}
}

// Create implements internal.BankRepository.
func (b *bank) Create(ctx context.Context, data *entity.Bank) error {
	bankRecord := record.BankEntityToRecord(data)

	_, err := b.db.NamedExecContext(ctx, query.QueryInsertBank, bankRecord)
	if err != nil {
		return err
	}

	return nil
}
