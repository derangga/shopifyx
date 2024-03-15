package entity

import "time"

type Payment struct {
	ID            int
	ProductID     int
	BankAccountID int
	ImageUrl      string
	Quantity      int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}
