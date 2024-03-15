package request

import (
	"time"

	"github.com/derangga/shopifyx/internal/entity"
)

type CreatePayment struct {
	ProductID       int    `param:"id" validate:"required"`
	BankAccountID   int    `json:"bankAccountID" validate:"required"`
	PaymentImageURL string `json:"paymentProofImageUrl" validate:"required,url"`
	Quantity        int    `json:"quantity" validate:"required,min=1"`
}

func (p *CreatePayment) ToEntityPayment() *entity.Payment {
	return &entity.Payment{
		ProductID:     p.ProductID,
		BankAccountID: p.BankAccountID,
		ImageUrl:      p.PaymentImageURL,
		Quantity:      p.Quantity,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     time.Time{},
	}
}
