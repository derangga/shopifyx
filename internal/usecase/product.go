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

func (uc *product) Update(ctx context.Context, data *entity.Product) (*entity.Product, error) {
	userID := pkgcontext.GetUserIDContext(ctx)
	data.UserID = userID

	product, err := uc.productRepo.Update(ctx, data)

	if _, ok := err.(errorpkg.ForbiddenAction); ok {
		log.Errorf("productUC.Update failed to uc.productRepo.Update: %w", err)
		return nil, errorpkg.NewCustomMessageError(err.Error(), http.StatusForbidden, err)
	}
	if _, ok := err.(errorpkg.RowNotFound); ok {
		log.Errorf("productUC.Update failed to uc.productRepo.Update: %w", err)
		return nil, errorpkg.NewCustomMessageError(err.Error(), http.StatusNotFound, err)
	}
	if err != nil {
		log.Errorf("productUC.Update failed to uc.productRepo.Update: %w", err)
		return nil, errorpkg.NewCustomError(http.StatusInternalServerError, err)
	}

	return product, nil
}

func (uc *product) Delete(ctx context.Context, data *entity.Product) error {
	userID := pkgcontext.GetUserIDContext(ctx)
	data.UserID = userID

	err := uc.productRepo.Delete(ctx, data)

	if _, ok := err.(errorpkg.ForbiddenAction); ok {
		log.Errorf("productUC.Delete failed to uc.productRepo.Delete: %w", err)
		return errorpkg.NewCustomMessageError(err.Error(), http.StatusForbidden, err)
	}
	if _, ok := err.(errorpkg.RowNotFound); ok {
		log.Errorf("productUC.Delete failed to uc.productRepo.Delete: %w", err)
		return errorpkg.NewCustomMessageError(err.Error(), http.StatusNotFound, err)
	}
	if err != nil {
		log.Errorf("productUC.Delete failed to uc.productRepo.Delete: %w", err)
		return errorpkg.NewCustomError(http.StatusInternalServerError, err)
	}

	return nil
}

func (uc *product) UpdateStock(ctx context.Context, data *entity.Product) error {
	userID := pkgcontext.GetUserIDContext(ctx)
	data.UserID = userID

	if userID != data.ID {
		return errorpkg.ForbiddenAction{
			Message: "You're not allowed to update this product",
		}
	}

	err := uc.productRepo.UpdateStock(ctx, data)

	if _, ok := err.(errorpkg.ForbiddenAction); ok {
		log.Errorf("productUC.UpdateStock failed to uc.productRepo.UpdateStock: %w", err)
		return errorpkg.NewCustomMessageError(err.Error(), http.StatusForbidden, err)
	}
	if _, ok := err.(errorpkg.RowNotFound); ok {
		log.Errorf("productUC.UpdateStock failed to uc.productRepo.UpdateStock: %w", err)
		return errorpkg.NewCustomMessageError(err.Error(), http.StatusNotFound, err)
	}
	if err != nil {
		log.Errorf("productUC.UpdateStock failed to uc.productRepo.UpdateStock: %w", err)
		return errorpkg.NewCustomError(http.StatusInternalServerError, err)
	}

	return nil
}

func (uc *product) Fetch(ctx context.Context, filter entity.ListFilter) (entity.ListResponse, error) {
	userID := pkgcontext.GetUserIDContext(ctx)
	filter.UserID = userID

	result, pagination, err := uc.productRepo.Fetch(ctx, filter)
	if err != nil {
		return entity.ListResponse{}, err
	}

	return entity.ListResponse{
		Data: result,
		Meta: entity.MetaTpl{
			Offset: filter.Page,
			Limit:  filter.Limit,
			Total:  pagination.Total,
		},
	}, nil
}
