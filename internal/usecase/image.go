package usecase

import (
	"bytes"
	"context"
	"net/http"

	"github.com/derangga/shopifyx/internal"
	"github.com/derangga/shopifyx/internal/config"
	"github.com/derangga/shopifyx/internal/constant"
	errorpkg "github.com/derangga/shopifyx/internal/pkg/error"
	"github.com/derangga/shopifyx/internal/pkg/helper"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
)

type imageUC struct {
	imageRepo internal.ImageRepository
	cfg       *config.BucketConfig
}

func NewImageUsecase(imageRepo internal.ImageRepository, cfg *config.BucketConfig) internal.ImageUsecase {
	return &imageUC{
		imageRepo: imageRepo,
		cfg:       cfg,
	}
}

// Upload implements internal.ImageUsecase.
func (uc *imageUC) Upload(ctx context.Context, req *bytes.Buffer) (string, error) {
	fileName := uuid.New().String() + constant.ExtensionJPEG
	filePath := helper.ConstructPath(constant.AWSPrefixPath, fileName)

	err := uc.imageRepo.Upload(ctx, uc.cfg.BucketName, filePath, req)
	if err != nil {
		log.Errorf("imageUC.Login failed to uc.imageRepo.Upload: %w", err)
		return "", errorpkg.NewCustomError(http.StatusInternalServerError, err)
	}

	return helper.ConstructPath(uc.cfg.BaseURL, filePath), nil
}
