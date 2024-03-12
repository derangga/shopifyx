package handler

import (
	"net/http"

	"github.com/derangga/shopifyx/internal"
	"github.com/derangga/shopifyx/internal/http/request"
	"github.com/derangga/shopifyx/internal/http/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// AuthHandler is a struct to handle http request in auth handler
type AuthHandler struct {
	authUC   internal.AuthUsecase
	validate *validator.Validate
}

// NewAuthHandler is a function to initialize auth handler
func NewAuthHandler(authUC internal.AuthUsecase, validate *validator.Validate) *AuthHandler {
	return &AuthHandler{
		authUC:   authUC,
		validate: validate,
	}
}

func (h *AuthHandler) Login(c echo.Context) error {
	return nil
}

func (h *AuthHandler) Register(c echo.Context) error {
	// bind request to struct
	var req request.Register
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: http.StatusText(http.StatusBadRequest),
		})
	}

	// validate request data
	err = h.validate.Struct(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: http.StatusText(http.StatusBadRequest),
		})
	}

	// proceed to usecase
	user, err := h.authUC.Register(c.Request().Context(), req.ToEntityUser())
	if err != nil {
		return NewCustomErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated, response.BaseResponse{
		Message: "User registered successfully",
		Data:    response.UserEntityToAuthResponse(user),
	})
}
