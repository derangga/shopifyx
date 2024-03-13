package handler

type Handlers struct {
	*AuthHandler
	*BankHandler
	*ProductHandler
}

func NewHandlers(
	authHandler *AuthHandler,
	bankHandler *BankHandler,
	productHandler *ProductHandler,
) *Handlers {
	return &Handlers{
		authHandler,
		bankHandler,
		productHandler,
	}
}
