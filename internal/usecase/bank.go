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
		log.Errorf("bankUC.Create failed to uc.bankRepo.Create: %w", err)
		return errorpkg.NewCustomError(http.StatusInternalServerError, err)
	}

	return nil
}
