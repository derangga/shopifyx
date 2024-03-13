package usecase

import (
	"context"
	"net/http"

	"github.com/derangga/shopifyx/internal"
	"github.com/derangga/shopifyx/internal/entity"
	pkgcontext "github.com/derangga/shopifyx/internal/pkg/context"
	errorpkg "github.com/derangga/shopifyx/internal/pkg/error"
	"github.com/labstack/gommon/log"
)

type product struct {
	productRepo internal.ProductRepository
	uow         internal.UnitOfWork
}

func NewProductUsecase(productRepo internal.ProductRepository, uow internal.UnitOfWork) internal.ProductUsecase {
	return &product{
		productRepo: productRepo,
		uow:         uow,
	}
}

// Create implements internal.ProductUsecase.
func (uc *product) Create(ctx context.Context, data *entity.Product) (*entity.Product, error) {
	userID := pkgcontext.GetUserIDContext(ctx)
	data.UserID = userID

	data, err := uc.productRepo.Create(ctx, data)
	if err != nil {
		log.Errorf("productUC.Create failed to uc.productRepo.Create: %w", err)
		return nil, errorpkg.NewCustomError(http.StatusInternalServerError, err)
	}

	return data, nil
}

func (uc *product) Update(ctx context.Context, id int, data *entity.Product) error {
	err := uc.productRepo.Update(ctx, id, data)
	if err != nil {
		log.Errorf("productUC.Update failed to uc.productRepo.Update: %d %w", id, err)
		return errorpkg.NewCustomError(http.StatusInternalServerError, err)
	}

	return nil
}
