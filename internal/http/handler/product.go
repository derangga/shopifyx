package handler

import (
	"net/http"
	"strconv"

	"github.com/derangga/shopifyx/internal"
	"github.com/derangga/shopifyx/internal/http/request"
	"github.com/derangga/shopifyx/internal/http/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// ProductHandler is a struct to handle http request in product handler
type ProductHandler struct {
	productUC internal.ProductUsecase
	validate  *validator.Validate
}

// NewProductHandler is a function to initialize product handler
func NewProductHandler(productUC internal.ProductUsecase, validate *validator.Validate) *ProductHandler {
	return &ProductHandler{
		productUC: productUC,
		validate:  validate,
	}
}

func (h *ProductHandler) Create(c echo.Context) error {
	// bind request to struct
	var req request.Product
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
	product, err := h.productUC.Create(c.Request().Context(), req.ToEntityProduct())
	if err != nil {
		return NewCustomErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated, response.BaseResponse{
		Message: http.StatusText(http.StatusCreated),
		Data:    response.ProductEntityToResponse(product),
	})
}

func (h *ProductHandler) Update(c echo.Context) error {
	var req request.ProductUpdate
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: http.StatusText(http.StatusBadRequest),
		})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: http.StatusText(http.StatusBadRequest),
		})
	}

	err = h.productUC.Update(c.Request().Context(), id, req.ToEntityProduct())
	if err != nil {
		return NewCustomErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Message: http.StatusText(http.StatusOK),
	})
}
