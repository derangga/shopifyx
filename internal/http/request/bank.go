package request

import "github.com/derangga/shopifyx/internal/entity"

type CreateBank struct {
	BankName          string `json:"bankName" validate:"required,min=5,max=15"`
	BankAccountName   string `json:"bankAccountName" validate:"required,min=5,max=15"`
	BankAccountNumber string `json:"bankAccountNumber" validate:"required,min=5,max=15"`
}

func (b *CreateBank) ToEntityBank() *entity.Bank {
	return &entity.Bank{
		BankName:          b.BankName,
		BankAccountName:   b.BankAccountName,
		BankAccountNumber: b.BankAccountNumber,
	}
}
