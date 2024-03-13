package repository

import (
	"context"

	"github.com/derangga/shopifyx/internal"
	"github.com/derangga/shopifyx/internal/entity"
	errorpkg "github.com/derangga/shopifyx/internal/pkg/error"
	"github.com/derangga/shopifyx/internal/repository/query"
	"github.com/derangga/shopifyx/internal/repository/record"
	"github.com/jmoiron/sqlx"
)

type product struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) internal.ProductRepository {
	return &product{
		db: db,
	}
}

// Get implements internal.UserRepository.
func (u *product) Get(ctx context.Context, id int) (*entity.Product, error) {
	var productRecord record.Product

	err := u.db.GetContext(ctx, &productRecord, query.ProductGetByID, id)
	if err != nil {
		return nil, err
	}

	return productRecord.ToEntity(), nil
}

// Create implements internal.ProductRepository.
func (u *product) Create(ctx context.Context, data *entity.Product) (*entity.Product, error) {
	return handleTransaction(ctx, u.db, func(ctx context.Context, tx *sqlx.Tx) (*entity.Product, error) {
		productRecord := record.ProductEntityToRecord(data)

		stmt, err := tx.PrepareNamedContext(ctx, query.ProductInsertQuery)
		if err != nil {
			return nil, err
		}

		row := stmt.QueryRowxContext(ctx, productRecord)
		if row.Err() != nil {
			return nil, row.Err()
		}

		err = row.Scan(&data.ID)
		if err != nil {
			return nil, err
		}

		return data, nil
	})
}

func (u *product) Update(ctx context.Context, id int, data *entity.Product) error {
	r := record.ProductEntityToRecord(data)
	r.ID = id

	res, err := u.db.ExecContext(ctx, query.ProductUpdate, r.ID, r.Name, r.Price, r.ImageURL, r.Condition, r.Tags, r.IsPurchaseable)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if rows == 0 {
		return errorpkg.RowNotFound{
			Message: "Product not found",
		}
	}

	return err
}
