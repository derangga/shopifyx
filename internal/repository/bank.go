package repository

import (
	"context"
	"net/http"

	"github.com/derangga/shopifyx/internal"
	"github.com/derangga/shopifyx/internal/entity"
	errorpkg "github.com/derangga/shopifyx/internal/pkg/error"
	"github.com/derangga/shopifyx/internal/repository/query"
	"github.com/derangga/shopifyx/internal/repository/record"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type bank struct {
	db *sqlx.DB
}

func NewBankRepository(db *sqlx.DB) internal.BankRepository {
	return &bank{
		db: db,
	}
}

// Fetch implements internal.BankRepository.
func (b *bank) Fetch(ctx context.Context, userID int) ([]entity.ListBank, error) {
	var result []entity.ListBank

	err := b.db.SelectContext(ctx, &result, query.QueryFetchBank, userID)
	if err != nil {
		log.Errorf("failed to fetch bank account: %w", err)
		return result, errorpkg.NewCustomMessageError("fatal query error", http.StatusInternalServerError, err)
	}

	return result, nil
}

// Create implements internal.BankRepository.
func (b *bank) Create(ctx context.Context, data *entity.Bank) error {
	bankRecord := record.BankEntityToRecord(data)

	_, err := b.db.NamedExecContext(ctx, query.QueryInsertBank, bankRecord)
	if err != nil {
		log.Errorf("failed to create bank account: %w", err)
		return errorpkg.NewCustomMessageError("fatal query error", http.StatusInternalServerError, err)
	}

	return nil
}

// Update implements internal.BankRepository.
func (b *bank) Update(ctx context.Context, data *entity.Bank) error {
	bankRecord := record.BankEntityToRecord(data)

	res, err := b.db.ExecContext(ctx, query.QueryUpdateBank, bankRecord.ID, bankRecord.UserID, bankRecord.BankName, bankRecord.BankAccountName, bankRecord.BankAccountNumber)
	if err != nil {
		log.Errorf("failed to update bank account: %w", err)
		return errorpkg.NewCustomError(http.StatusInternalServerError, err)
	}

	row, err := res.RowsAffected()
	if err != nil {
		log.Errorf("failed retrieve affected rows: %w", err)
		return errorpkg.NewCustomError(http.StatusInternalServerError, err)
	}
	if row == 0 {
		return errorpkg.NewCustomMessageError("bank account not found", http.StatusNotFound, err)
	}

	return nil
}

// SoftDelete implements internal.BankRepository.
func (b *bank) SoftDelete(ctx context.Context, data *entity.Bank) error {
	bankRecord := record.BankEntityToRecord(data)

	res, err := b.db.ExecContext(ctx, query.QuerySoftDeleteBank, bankRecord.ID, bankRecord.UserID)
	if err != nil {
		log.Errorf("failed to delete bank account: %w", err)
		return errorpkg.NewCustomError(http.StatusInternalServerError, err)
	}

	row, err := res.RowsAffected()
	if err != nil {
		log.Errorf("failed retrieve affected rows: %w", err)
		return errorpkg.NewCustomError(http.StatusInternalServerError, err)
	}
	if row == 0 {
		return errorpkg.NewCustomMessageError("bank account not found", http.StatusNotFound, err)
	}

	return nil
}

// Get implements internal.BankRepository.
func (b *bank) Get(ctx context.Context, id int) (*entity.Bank, error) {
	var bankRecord record.Bank

	err := b.db.GetContext(ctx, &bankRecord, query.QueryBankGetByID, id)
	if err != nil {
		return nil, err
	}

	return bankRecord.ToEntity(), nil
}
