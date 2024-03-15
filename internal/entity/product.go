package entity

import (
	"time"
)

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

type ListProduct struct {
	ID             int      `json:"productId" db:"id"`
	Name           string   `json:"name" db:"name"`
	Price          int      `json:"price" db:"price"`
	ImageUrl       string   `json:"imageUrl" db:"image_url"`
	Stock          int      `json:"stock" db:"stock"`
	Condition      string   `json:"condition" db:"condition"`
	Tags           []string `json:"tags" db:"tags"`
	IsPurchaseable bool     `json:"isPurchaseable" db:"is_purchaseable"`
	PurchaseCount  int      `json:"purchaseCount" db:"purchase_count"`
	Total          int      `json:"-" db:"total"`
}

type ListFilter struct {
	UserOnly       bool
	Tags           []string
	Condition      string
	ShowEmptyStock bool
	MaxPrice       int
	MinPrice       int
	SortBy         string
	OrderBy        string
	Search         string
	Page           int
	Limit          int

	UserID int
}

type MetaTpl struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type ListResponse struct {
	Data []ListProduct `json:"products"`
	Meta MetaTpl       `json:"meta"`
}
