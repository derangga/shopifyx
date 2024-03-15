package entity

import "time"

type Bank struct {
	ID                int
	UserID            int
	BankName          string
	BankAccountName   string
	BankAccountNumber string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         time.Time
}

type ListBank struct {
	ID                int    `json:"bankId" db:"id"`
	BankName          string `json:"bankName" db:"bank_name"`
	BankAccountName   string `json:"bankAccountName" db:"bank_account_name"`
	BankAccountNumber string `json:"bankAccountNumber" db:"bank_account_number"`
}
