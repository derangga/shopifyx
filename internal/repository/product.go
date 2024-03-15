package repository

import (
	"context"
	"time"

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

func (u *product) Update(ctx context.Context, data *entity.Product) error {
	r := record.ProductEntityToRecord(data)

	res, err := u.db.ExecContext(
		ctx,
		query.ProductUpdate,
		r.ID, r.UserID,
		r.Name, r.Price,
		r.ImageURL,
		r.Condition,
		r.Tags,
		r.IsPurchaseable,
		time.Now(),
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errorpkg.RowNotFound{
			Message: "Product not found",
		}
	}

	return nil
}

func (u *product) Delete(ctx context.Context, data *entity.Product) error {
	res, err := u.db.ExecContext(ctx, query.ProductDelete, data.ID, data.UserID, time.Now())
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errorpkg.RowNotFound{
			Message: "No matching product deleted",
		}
	}

	return nil
}

func (u *product) UpdateStock(ctx context.Context, data *entity.Product) error {
	res, err := u.db.ExecContext(ctx, query.ProductStockUpdate, data.ID, data.UserID, data.Stock)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errorpkg.RowNotFound{
			Message: "No matching product, stock is not updated",
		}
	}

	return nil
}
