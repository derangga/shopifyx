package internal

import (
	"bytes"
	"context"

	"github.com/derangga/shopifyx/internal/entity"
)

type UserRepository interface {
	Get(ctx context.Context, id int) (*entity.User, error)
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	Create(ctx context.Context, req *entity.User) (*entity.User, error)
}

type BankRepository interface {
	Create(ctx context.Context, req *entity.Bank) error
	Update(ctx context.Context, req *entity.Bank) error
	SoftDelete(ctx context.Context, req *entity.Bank) error
	Fetch(ctx context.Context, userID int) ([]entity.ListBank, error)
	Get(ctx context.Context, id int) (*entity.Bank, error)
}

type ProductRepository interface {
	Get(ctx context.Context, id int) (*entity.Product, error)
	Create(ctx context.Context, req *entity.Product) (*entity.Product, error)
	GetDetailedByID(ctx context.Context, id int, userId int) (*entity.ProductDetail, error)
	Update(ctx context.Context, req *entity.Product) (*entity.Product, error)
	Delete(ctx context.Context, req *entity.Product) error
	UpdateStock(ctx context.Context, req *entity.Product) error
	Fetch(ctx context.Context, filter entity.ListFilter) ([]entity.ListProduct, *entity.MetaTpl, error)
}

type ImageRepository interface {
	Upload(ctx context.Context, bucket string, key string, file *bytes.Buffer) error
}

type PaymentRepository interface {
	Create(ctx context.Context, req *entity.Payment) error
}

type UnitOfWork interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
