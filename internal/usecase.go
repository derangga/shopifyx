package internal

import (
	"bytes"
	"context"

	"github.com/derangga/shopifyx/internal/entity"
)

type AuthUsecase interface {
	Register(ctx context.Context, req *entity.User) (*entity.User, error)
	Login(ctx context.Context, req *entity.User) (*entity.User, error)
}

type BankUsecase interface {
	Create(ctx context.Context, req *entity.Bank) error
	Update(ctx context.Context, req *entity.Bank) error
	Delete(ctx context.Context, id int) error
	Fetch(ctx context.Context) ([]entity.ListBank, error)
}

type ImageUsecase interface {
	Upload(ctx context.Context, req *bytes.Buffer) (string, error)
}

type ProductUsecase interface {
	Create(ctx context.Context, req *entity.Product) (*entity.Product, error)
	Update(ctx context.Context, req *entity.Product) (*entity.Product, error)
	Delete(ctx context.Context, req *entity.Product) error
	UpdateStock(ctx context.Context, req *entity.Product) error
	Fetch(ctx context.Context, filter entity.ListFilter) (entity.ListResponse, error)
}

type PaymentUsecase interface {
	Create(ctx context.Context, req *entity.Payment) error
}
