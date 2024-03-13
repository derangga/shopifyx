package internal

import (
	"context"

	"github.com/derangga/shopifyx/internal/entity"
)

// Usecases is list of available usecases
type Usecases struct {
	AuthUsecase AuthUsecase
	BankUsecase BankUsecase
}

type AuthUsecase interface {
	Register(ctx context.Context, req *entity.User) (*entity.User, error)
	Login(ctx context.Context, req *entity.User) (*entity.User, error)
}

type BankUsecase interface {
	Create(ctx context.Context, req *entity.Bank) error
}

type ProductUsecase interface {
	Create(ctx context.Context, req *entity.Product) (*entity.Product, error)
}
