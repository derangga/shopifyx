package usecase

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/derangga/shopifyx/internal"
	"github.com/derangga/shopifyx/internal/config"
	"github.com/derangga/shopifyx/internal/entity"
	errorpkg "github.com/derangga/shopifyx/internal/pkg/error"
	"github.com/derangga/shopifyx/internal/pkg/helper"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type auth struct {
	userRepo internal.UserRepository
	uow      internal.UnitOfWork
	cfg      *config.AuthConfig
}

func NewAuthUsecase(userRepo internal.UserRepository, uow internal.UnitOfWork, cfg *config.AuthConfig) internal.AuthUsecase {
	return &auth{
		userRepo: userRepo,
		uow:      uow,
		cfg:      cfg,
	}
}

// Login implements internal.AuthUsecase.
func (uc *auth) Login(ctx context.Context, req *entity.User) (*entity.User, error) {
	// get user by username
	user, err := uc.userRepo.GetByUsername(ctx, req.Username)
	if err != nil && !helper.IsSQLErrNotFound(err) {
		log.Errorf("authUC.Login failed to uc.userRepo.GetByUsername: %w", err)
		return nil, errorpkg.NewCustomError(http.StatusInternalServerError, err)
	}

	// return error if user already exists
	if user == nil {
		return nil, errorpkg.NewCustomError(http.StatusNotFound, err)
	}

	// return error if password doesnt match
	if !uc.validateHash(user.Password, req.Password) {
		return nil, errorpkg.NewCustomError(http.StatusBadRequest, err)
	}

	// construct jwt access token
	user.AccessToken, err = uc.constructJWT(user)
	if err != nil {
		log.Errorf("authUC.Register failed to uc.constructJWT: %w", err)
		return nil, errorpkg.NewCustomError(http.StatusInternalServerError, err)
	}

	return user, nil
}

// Register implements internal.AuthUsecase.
func (uc *auth) Register(ctx context.Context, req *entity.User) (*entity.User, error) {
	// get user by username
	user, err := uc.userRepo.GetByUsername(ctx, req.Username)
	if err != nil && !helper.IsSQLErrNotFound(err) {
		log.Errorf("authUC.Register failed to uc.userRepo.GetByUsername: %w", err)
		return nil, errorpkg.NewCustomError(http.StatusInternalServerError, err)
	}

	// return error if user already exists
	if user != nil {
		return nil, errorpkg.NewCustomError(http.StatusConflict, err)
	}

	// hash password
	hashed, err := uc.generateHash(req.Password)
	if err != nil {
		log.Errorf("authUC.Register failed to uc.generateHash: %w", err)
		return nil, errorpkg.NewCustomError(http.StatusInternalServerError, err)
	}

	// construct user data
	req.Password = hashed
	req.CreatedAt = time.Now()

	// save user to db
	user, err = uc.userRepo.Create(ctx, req)
	if err != nil {
		log.Errorf("authUC.Register failed to uc.userRepo.Create: %w", err)
		return nil, errorpkg.NewCustomError(http.StatusInternalServerError, err)
	}

	// construct jwt access token
	user.AccessToken, err = uc.constructJWT(user)
	if err != nil {
		log.Errorf("authUC.Register failed to uc.constructJWT: %w", err)
		return nil, errorpkg.NewCustomError(http.StatusInternalServerError, err)
	}

	return user, nil
}

// generateHash implements internal.AuthUsecase.
func (uc *auth) generateHash(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), uc.cfg.BcryptSalt)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

// validateHash implements internal.AuthUsecase.
func (uc *auth) validateHash(hashed string, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}

// constructJWT implements internal.AuthUsecase.
func (uc *auth) constructJWT(req *entity.User) (string, error) {
	now := time.Now()
	claims := jwt.RegisteredClaims{
		Subject:   strconv.Itoa(req.ID),
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(uc.cfg.JWTValidDuration)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(uc.cfg.JWTSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
