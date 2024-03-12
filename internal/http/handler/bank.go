package handler

import (
	"net/http"

	"github.com/derangga/shopifyx/internal"
	"github.com/derangga/shopifyx/internal/http/request"
	"github.com/derangga/shopifyx/internal/http/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// BankHandler is a struct to handle http request in bank handler
type BankHandler struct {
	bankUC   internal.BankUsecase
	validate *validator.Validate
}

// NewBankHandler is a function to initialize bank handler
func NewBankHandler(bankUC internal.BankUsecase, validate *validator.Validate) *BankHandler {
	return &BankHandler{
		bankUC:   bankUC,
		validate: validate,
	}
}

func (h *BankHandler) Create(c echo.Context) error {
	// bind request to struct
	var req request.CreateBank
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
			Message: err.Error(),
		})
	}

	// proceed to usecase
	err = h.bankUC.Create(c.Request().Context(), req.ToEntityBank())
	if err != nil {
		return NewCustomErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated, response.BaseResponse{
		Message: "bank account created successfully",
		Data:    nil,
	})
}
