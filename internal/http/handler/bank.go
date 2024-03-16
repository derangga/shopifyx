package handler

import (
	"fmt"
	"net/http"
	"strconv"

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
			Message: err.Error(),
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

	return c.JSON(http.StatusOK, response.BaseResponse{
		Message: "bank account created successfully",
	})
}

func (h *BankHandler) Update(c echo.Context) error {
	var req request.UpdateBank

	id, err := strconv.Atoi(c.Param("bankAccountId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: fmt.Sprintf("%s should be integer, got error: %v", "bankAccountId", err),
		})
	}

	req.ID = id
	err = c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: err.Error(),
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
	err = h.bankUC.Update(c.Request().Context(), req.ToEntityBank())
	if err != nil {
		return NewCustomErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Message: "bank account updated successfully",
		Data:    nil,
	})
}

func (h *BankHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("bankAccountId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: fmt.Sprintf("%s should be integer, got error: %v", "bankAccountId", err),
		})
	}

	// proceed to usecase
	err = h.bankUC.Delete(c.Request().Context(), id)
	if err != nil {
		return NewCustomErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Message: "bank account deleted successfully",
		Data:    nil,
	})
}

func (h *BankHandler) Fetch(c echo.Context) error {
	// proceed to usecase
	res, err := h.bankUC.Fetch(c.Request().Context())
	if err != nil {
		return NewCustomErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Message: "success",
		Data:    res,
	})
}
