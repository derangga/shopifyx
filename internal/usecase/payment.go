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

type payment struct {
	paymentRepo internal.PaymentRepository
	productRepo internal.ProductRepository
	bankRepo    internal.BankRepository
	uow         internal.UnitOfWork
}

func NewPaymentUsecase(paymentRepo internal.PaymentRepository, productRepo internal.ProductRepository, bankRepo internal.BankRepository, uow internal.UnitOfWork) internal.PaymentUsecase {
	return &payment{
		paymentRepo: paymentRepo,
		productRepo: productRepo,
		bankRepo:    bankRepo,
		uow:         uow,
	}
}

func (uc *payment) Create(ctx context.Context, req *entity.Payment) error {
	userID := pkgcontext.GetUserIDContext(ctx)

	bank, err := uc.bankRepo.Get(ctx, req.BankAccountID)
	if err != nil {
		log.Error(err)
		return err
	}

	if bank.ID != userID {
		return errorpkg.ForbiddenAction{
			Message: "You're not allowed to update this product",
		}
	}

	product, err := uc.productRepo.Get(ctx, req.ProductID)
	if err != nil {
		log.Error(err)
		return err
	}

	if product.Stock < req.Quantity {
		return errorpkg.NewCustomMessageError("insufficient stock", http.StatusBadRequest, nil)
	}

	err = uc.uow.WithTransaction(ctx, func(ctx context.Context) error {
		product.Stock -= req.Quantity
		err = uc.productRepo.UpdateStock(ctx, product)
		if err != nil {
			log.Error(err)
			return err
		}

		product.PurchaseCount++
		err = uc.productRepo.UpdatePurchaseCount(ctx, product)
		if err != nil {
			log.Error(err)
			return err
		}

		err = uc.paymentRepo.Create(ctx, req)
		if err != nil {
			log.Error(err)
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
