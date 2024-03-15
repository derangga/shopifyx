package handler

import (
	"net/http"

	"github.com/derangga/shopifyx/internal"
	"github.com/derangga/shopifyx/internal/http/request"
	"github.com/derangga/shopifyx/internal/http/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	paymentUC internal.PaymentUsecase
	validate  *validator.Validate
}

func NewPaymentHandler(paymentUC internal.PaymentUsecase, validate *validator.Validate) *PaymentHandler {
	return &PaymentHandler{
		paymentUC: paymentUC,
		validate:  validate,
	}
}

func (h *PaymentHandler) Create(c echo.Context) error {

	var req request.CreatePayment

	//bind
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: err.Error(),
		})
	}

	//validate
	err = h.validate.Struct(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: err.Error(),
		})
	}

	//proceed to usecase
	err = h.paymentUC.Create(c.Request().Context(), req.ToEntityPayment())
	if err != nil {
		return NewCustomErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated, response.BaseResponse{
		Message: "payment process succesfully",
		Data:    nil,
	})
}
