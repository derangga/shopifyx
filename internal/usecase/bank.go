package usecase

import (
	"context"

	"github.com/derangga/shopifyx/internal"
	"github.com/derangga/shopifyx/internal/entity"
	pkgcontext "github.com/derangga/shopifyx/internal/pkg/context"
)

type bank struct {
	bankRepo internal.BankRepository
}

func NewBankUsecase(bankRepo internal.BankRepository) internal.BankUsecase {
	return &bank{
		bankRepo: bankRepo,
	}
}

// Create implements internal.BankUsecase.
func (uc *bank) Create(ctx context.Context, req *entity.Bank) error {
	req.UserID = pkgcontext.GetUserIDContext(ctx)

	err := uc.bankRepo.Create(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

// Update implements internal.BankUsecase.
func (uc *bank) Update(ctx context.Context, req *entity.Bank) error {
	req.UserID = pkgcontext.GetUserIDContext(ctx)

	err := uc.bankRepo.Update(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

// Delete implements internal.BankUsecase.
func (uc *bank) Delete(ctx context.Context, id int) error {
	userID := pkgcontext.GetUserIDContext(ctx)

	err := uc.bankRepo.SoftDelete(ctx, &entity.Bank{ID: id, UserID: userID})
	if err != nil {
		return err
	}

	return nil
}
