package repository

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/derangga/shopifyx/internal"
	"github.com/derangga/shopifyx/internal/entity"
	errorpkg "github.com/derangga/shopifyx/internal/pkg/error"
	"github.com/derangga/shopifyx/internal/repository/query"
	"github.com/derangga/shopifyx/internal/repository/record"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	"github.com/lib/pq"
)

type product struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) internal.ProductRepository {
	return &product{
		db: db,
	}
}

func (u *product) validateProductBeforeModified(ctx context.Context, id int, userId int) (*entity.Product, error) {
	product, err := u.Get(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, errorpkg.RowNotFound{
				Message: "Product not found",
			}
		}
		return nil, err
	}
	if product == nil {
		return nil, errorpkg.RowNotFound{
			Message: "Product not found",
		}
	}
	if product.UserID != userId {
		return nil, errorpkg.ForbiddenAction{
			Message: "You're not allowed to update this product",
		}
	}
	return product, nil
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

func (u *product) GetDetailedByID(ctx context.Context, id int, userId int) (*entity.ProductDetail, error) {
	var productDetailRecord record.ProductDetail
	var productSoldTotal int

	err := u.db.GetContext(ctx, &productDetailRecord, query.ProductDetailByID, id)

	if err != nil {
		return nil, err
	}

	err = u.db.GetContext(ctx, &productSoldTotal, query.ProductUserSoldTotal, userId)
	if err != nil {
		return nil, err
	}

	entity := productDetailRecord.ToEntity()
	entity.Seller.ProductSoldTotal = productSoldTotal

	return entity, nil
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

func (u *product) Update(ctx context.Context, data *entity.Product) (*entity.Product, error) {

	product, err := u.validateProductBeforeModified(ctx, data.ID, data.UserID)
	if err != nil {
		return nil, err
	}

	r := record.ProductEntityToRecord(data)

	updatedAt := time.Now()
	_, err = u.db.ExecContext(
		ctx,
		query.ProductUpdate,
		r.ID, r.UserID,
		r.Name, r.Price,
		r.ImageURL,
		r.Condition,
		r.Tags,
		r.IsPurchaseable,
		updatedAt,
	)

	if err != nil {
		return nil, err
	}
	product.UpdatedAt = updatedAt

	return product, nil
}

func (u *product) Delete(ctx context.Context, data *entity.Product) error {
	_, err := u.validateProductBeforeModified(ctx, data.ID, data.UserID)
	if err != nil {
		return err
	}

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
	//modify it to allow reuse by payment
	_, err := handleTransaction(ctx, u.db, func(ctx context.Context, tx *sqlx.Tx) (*entity.Product, error) {
		_, err := tx.ExecContext(ctx, query.ProductStockUpdate, data.ID, data.UserID, data.Stock)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	return err
}

func (u *product) Fetch(ctx context.Context, filter entity.ListFilter) ([]entity.ListProduct, *entity.MetaTpl, error) {
	var conditions []string
	var values []interface{}

	query := `SELECT
		COUNT(p.id) OVER() AS total,
		p.id,
		p.name,
		p.price,
		p.image_url,
		p.stock,
		p."condition",
		p.tags,
		p.is_purchaseable,
		p.purchase_count
	FROM product p`

	// only retrieve product active
	conditions = append(conditions, "p.deleted_at is null")

	// filter tags
	orCond := []string{}
	for _, t := range filter.Tags {
		orCond = append(orCond, "p.tags @> ?")
		values = append(values, fmt.Sprintf("{%s}", t))
	}
	if len(orCond) > 1 {
		conditions = append(conditions, fmt.Sprintf("(%s)", strings.Join(orCond, " OR ")))
	}
	if len(orCond) == 1 {
		conditions = append(conditions, orCond...)
	}

	// filter condition
	if filter.Condition != "" {
		conditions = append(conditions, "p.condition = ?")
		values = append(values, filter.Condition)
	}

	// filter empty stock
	if filter.ShowEmptyStock {
		conditions = append(conditions, "p.stock >= 0")
	} else {
		conditions = append(conditions, "p.stock > 0")
	}

	// filter min and max price
	if filter.MinPrice > 0 && filter.MaxPrice > 0 {
		conditions = append(conditions, "p.price BETWEEN ? AND ?")
		values = append(values, filter.MinPrice)
		values = append(values, filter.MaxPrice)
	}

	// filter search
	if len(filter.Search) > 0 {
		conditions = append(conditions, "p.name ILIKE '%' || ? || '%'")
		values = append(values, filter.Search)
	}

	// filter created by
	if filter.UserOnly {
		conditions = append(conditions, "p.user_id = ?")
		values = append(values, filter.UserID)
	}

	// Build condition
	if len(conditions) > 0 && len(values) > 0 {
		query = fmt.Sprintf("%s WHERE %s", query, strings.Join(conditions, " AND "))
	}

	// build grouping
	query = fmt.Sprintf("%s GROUP BY p.id, p.created_at", query)

	// build sorting
	query = fmt.Sprintf("%s ORDER BY p.created_at DESC", query)
	if len(filter.OrderBy) > 0 && len(filter.SortBy) > 0 {
		query = fmt.Sprintf("%s, %s %s", query, filter.SortBy, filter.OrderBy)
	}

	// Build pagination
	limit := 10
	if filter.Limit > 0 {
		limit = filter.Limit
	}
	page := 1
	if filter.Page > 0 {
		page = filter.Page
	}
	offset := limit * (page - 1)

	query = fmt.Sprintf("%s LIMIT ? OFFSET ? ", query)
	values = append(values, limit)
	values = append(values, offset)

	query = u.db.Rebind(query)

	result := make([]entity.ListProduct, 0)
	pagination := new(entity.MetaTpl)

	rows, err := u.db.Query(query, values...)
	if err != nil {
		log.Errorf("failed to query list products: %w", err)
		return nil, nil, errorpkg.NewCustomMessageError("fatal query error", http.StatusInternalServerError, err)
	}

	defer rows.Close()
	for rows.Next() {
		d := entity.ListProduct{}
		err := rows.Scan(
			&d.Total,
			&d.ID,
			&d.Name,
			&d.Price,
			&d.ImageUrl,
			&d.Stock,
			&d.Condition,
			pq.Array(&d.Tags),
			&d.IsPurchaseable,
			&d.PurchaseCount,
		)
		if err != nil {
			log.Errorf("failed to query list products: %w", err)
			return nil, nil, errorpkg.NewCustomMessageError("fatal query error", http.StatusInternalServerError, err)
		}
		result = append(result, d)
	}

	err = rows.Err()
	if err != nil {
		log.Errorf("failed to query list products: %w", err)
		return nil, nil, errorpkg.NewCustomMessageError("fatal query error", http.StatusInternalServerError, err)
	}

	if len(result) == 0 {
		pagination.Total = 0
		return result, pagination, nil
	}

	pagination.Offset = page
	pagination.Limit = limit
	pagination.Total = result[0].Total

	return result, pagination, nil
}
