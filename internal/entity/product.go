package entity

import "time"

type Product struct {
	ID             int
	UserID         int
	Name           string
	Price          int
	ImageURL       string
	Stock          int
	Condition      string
	Tags           []string
	IsPurchaseable bool
	PurchaseCount  int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}
