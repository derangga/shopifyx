package app

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/derangga/shopifyx/internal/config"
	"github.com/derangga/shopifyx/internal/http"
	"github.com/derangga/shopifyx/internal/http/handler"
	"github.com/derangga/shopifyx/internal/http/middleware"
	"github.com/derangga/shopifyx/internal/pkg/database"
	"github.com/derangga/shopifyx/internal/repository"
	"github.com/derangga/shopifyx/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

var (
	configSet = wire.NewSet(
		provideAuthConfig,
		provideDBConfig,
		provideBucketConfig,
	)

	repositoriesSet = wire.NewSet(
		repository.NewUserRepository,
		repository.NewProductRepository,
		repository.NewBankRepository,
		repository.NewImageRepository,
		repository.NewUnitOfWork,
	)

	usecasesSet = wire.NewSet(
		usecase.NewAuthUsecase,
		usecase.NewBankUsecase,
		usecase.NewProductUsecase,
		usecase.NewImageUsecase,
	)

	handlerSet = wire.NewSet(
		handler.NewAuthHandler,
		handler.NewBankHandler,
		handler.NewProductHandler,
		handler.NewImageHandler,
	)

	middlewareSet = wire.NewSet(
		provideJWTAuth,
	)

	dependencySet = wire.NewSet(
		provideDB,
		provideValidator,
		provideS3Client,
	)

	allSet = wire.NewSet(
		configSet,
		dependencySet,
		repositoriesSet,
		usecasesSet,
		handlerSet,
		middlewareSet,
		handler.NewHandlers,
		http.NewHttpServer,
	)
)

func provideAuthConfig(cfg *config.Config) *config.AuthConfig {
	return &cfg.Auth
}

func provideDBConfig(cfg *config.Config) *config.DatabaseConfig {
	return &cfg.Database
}

func provideBucketConfig(cfg *config.Config) *config.BucketConfig {
	if cfg.Bucket.BaseURL == "" {
		cfg.Bucket.BaseURL = cfg.Bucket.ConstructURL()
	}

	return &cfg.Bucket
}

func provideDB(cfg *config.DatabaseConfig) *sqlx.DB {
	return database.NewPostgresDatabase(cfg)
}

func provideValidator() *validator.Validate {
	return validator.New()
}

func provideJWTAuth(cfg *config.AuthConfig) *middleware.JWTAuth {
	return middleware.ProvideJWTAuth(cfg.JWTSecret)
}

func provideS3Client(cfg *config.BucketConfig) *s3.S3 {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(cfg.Region),
		Credentials: credentials.NewStaticCredentials(
			cfg.ID,
			cfg.Secret,
			"",
		),
	})
	if err != nil {
		log.Fatalf("error initiate aws session: %w", err)
	}

	return s3.New(sess)
}
