package record

import (
	"database/sql"
	"time"

	"github.com/derangga/shopifyx/internal/entity"
	"github.com/derangga/shopifyx/internal/pkg/helper"
)

type Bank struct {
	ID                int          `db:"id"`
	BankName          string       `db:"bank_name"`
	BankAccountName   string       `db:"bank_account_name"`
	BankAccountNumber string       `db:"bank_account_number"`
	CreatedAt         time.Time    `db:"created_at"`
	UpdatedAt         sql.NullTime `db:"updated_at"`
	DeletedAt         sql.NullTime `db:"deleted_at"`
}

func (r *Bank) ToEntity() *entity.Bank {
	return &entity.Bank{
		ID:                r.ID,
		BankName:          r.BankName,
		BankAccountName:   r.BankAccountName,
		BankAccountNumber: r.BankAccountNumber,
		CreatedAt:         r.CreatedAt,
		UpdatedAt:         r.UpdatedAt.Time,
		DeletedAt:         r.DeletedAt.Time,
	}
}

func BankEntityToRecord(req *entity.Bank) *Bank {
	return &Bank{
		ID:                req.ID,
		BankName:          req.BankName,
		BankAccountName:   req.BankAccountName,
		BankAccountNumber: req.BankAccountNumber,
		CreatedAt:         req.CreatedAt,
		UpdatedAt:         helper.NullTime(req.UpdatedAt),
		DeletedAt:         helper.NullTime(req.DeletedAt),
	}
}
