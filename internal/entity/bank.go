package entity

import "time"

type Bank struct {
	ID                int
	BankName          string
	BankAccountName   string
	BankAccountNumber string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         time.Time
}
