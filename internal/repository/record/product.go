package record

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/derangga/shopifyx/internal/entity"
	"github.com/derangga/shopifyx/internal/pkg/helper"
)

type Product struct {
	ID             int          `db:"id"`
	UserID         int          `db:"user_id"`
	Name           string       `db:"name"`
	Price          int          `db:"price"`
	ImageURL       string       `db:"image_url"`
	Stock          int          `db:"stock"`
	Condition      string       `db:"condition"`
	Tags           string       `db:"tags"`
	IsPurchaseable bool         `db:"is_purchaseable"`
	PurchaseCount  int          `db:"purchase_count"`
	CreatedAt      time.Time    `db:"created_at"`
	UpdatedAt      sql.NullTime `db:"updated_at"`
	DeletedAt      sql.NullTime `db:"deleted_at"`
}

func (r *Product) ToEntity() *entity.Product {
	return &entity.Product{
		ID:        r.ID,
		UserID:    r.UserID,
		Name:      r.Name,
		Price:     r.Price,
		ImageURL:  r.ImageURL,
		Stock:     r.Stock,
		Condition: r.Condition,
		Tags: func() []string {
			if len(r.Tags) <= 2 {
				return nil
			}

			return strings.Split(r.Tags[1:len(r.Tags)-1], ",")
		}(),
		IsPurchaseable: r.IsPurchaseable,
		PurchaseCount:  r.PurchaseCount,
		CreatedAt:      r.CreatedAt,
		UpdatedAt:      r.UpdatedAt.Time,
		DeletedAt:      r.DeletedAt.Time,
	}
}

func ProductEntityToRecord(req *entity.Product) *Product {
	return &Product{
		ID:             req.ID,
		UserID:         req.UserID,
		Name:           req.Name,
		Price:          req.Price,
		ImageURL:       req.ImageURL,
		Stock:          req.Stock,
		Condition:      req.Condition,
		Tags:           fmt.Sprintf("{%s}", strings.Join(req.Tags, ",")),
		IsPurchaseable: req.IsPurchaseable,
		PurchaseCount:  req.PurchaseCount,
		CreatedAt:      req.CreatedAt,
		UpdatedAt:      helper.NullTime(req.UpdatedAt),
		DeletedAt:      helper.NullTime(req.DeletedAt),
	}
}
