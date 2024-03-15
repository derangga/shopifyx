package record

import (
	"database/sql"
	"time"

	"github.com/derangga/shopifyx/internal/entity"
	"github.com/derangga/shopifyx/internal/pkg/helper"
)

type Payment struct {
	ID            int          `db:"id"`
	ProductID     int          `db:"product_id"`
	BankAccountID int          `db:"bank_account_id"`
	ImageURL      string       `db:"payment_proof_image_url"`
	Quantity      int          `db:"quantity"`
	CreatedAt     time.Time    `db:"created_at"`
	UpdatedAt     sql.NullTime `db:"updated_at"`
	DeletedAt     sql.NullTime `db:"deleted_at"`
}

func PaymentEntityToRecord(req *entity.Payment) *Payment {
	return &Payment{
		ID:            req.ID,
		ProductID:     req.ProductID,
		BankAccountID: req.BankAccountID,
		ImageURL:      req.ImageUrl,
		Quantity:      req.Quantity,
		CreatedAt:     req.CreatedAt,
		UpdatedAt:     helper.NullTime(req.UpdatedAt),
		DeletedAt:     helper.NullTime(req.DeletedAt),
	}
}
