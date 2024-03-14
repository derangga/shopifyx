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

	if _, ok := err.(errorpkg.RowNotFound); ok {
		log.Errorf("productUC.Update failed to uc.productRepo.Update: %w", err)
		return errorpkg.NewCustomMessageError(err.Error(), http.StatusNotFound, err)
	} else if err != nil {
		log.Errorf("productUC.Update failed to uc.productRepo.Update: %w", err)
		return errorpkg.NewCustomError(http.StatusInternalServerError, err)
	} else {
		return nil
	}
}

func (uc *product) Delete(ctx context.Context, id int) error {
	err := uc.productRepo.Delete(ctx, id)

	if _, ok := err.(errorpkg.RowNotFound); ok {
		log.Errorf("productUC.Delete failed to uc.productRepo.Delete: %w", err)
		return errorpkg.NewCustomMessageError(err.Error(), http.StatusNotFound, err)
	} else if err != nil {
		log.Errorf("productUC.Delete failed to uc.productRepo.Delete: %w", err)
		return errorpkg.NewCustomError(http.StatusInternalServerError, err)
	} else {
		return nil
	}
}

func (uc *product) UpdateStock(ctx context.Context, id int, stock int) error {
	err := uc.productRepo.UpdateStock(ctx, id, stock)

	if _, ok := err.(errorpkg.RowNotFound); ok {
		log.Errorf("productUC.UpdateStock failed to uc.productRepo.UpdateStock: %w", err)
		return errorpkg.NewCustomMessageError(err.Error(), http.StatusNotFound, err)
	} else if err != nil {
		log.Errorf("productUC.UpdateStock failed to uc.productRepo.UpdateStock: %w", err)
		return errorpkg.NewCustomError(http.StatusInternalServerError, err)
	} else {
		return nil
	}
}
