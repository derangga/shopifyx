package handler

type Handlers struct {
	*AuthHandler
	*BankHandler
	*ProductHandler
	*ImageHandler
	*PaymentHandler
}

func NewHandlers(
	authHandler *AuthHandler,
	bankHandler *BankHandler,
	productHandler *ProductHandler,
	imageHandler *ImageHandler,
	paymentHandler *PaymentHandler,
) *Handlers {
	return &Handlers{
		authHandler,
		bankHandler,
		productHandler,
		imageHandler,
		paymentHandler,
	}
}
