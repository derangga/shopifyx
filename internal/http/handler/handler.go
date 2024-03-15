package handler

type Handlers struct {
	*AuthHandler
	*BankHandler
	*ProductHandler
	*ImageHandler
}

func NewHandlers(
	authHandler *AuthHandler,
	bankHandler *BankHandler,
	productHandler *ProductHandler,
	imageHandler *ImageHandler,
) *Handlers {
	return &Handlers{
		authHandler,
		bankHandler,
		productHandler,
		imageHandler,
	}
}
