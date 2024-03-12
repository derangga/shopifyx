package internal

import (
	"context"

	"github.com/derangga/shopifyx/internal/entity"
)

// Repositories is list of available repository
type Repositories struct {
	UserRepository UserRepository
	UOW            UnitOfWork
}

type UserRepository interface {
	Get(ctx context.Context, id int) (*entity.User, error)
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	Create(ctx context.Context, req *entity.User) (*entity.User, error)
}

type UnitOfWork interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}