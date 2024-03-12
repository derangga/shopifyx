package handler

type Handlers struct {
	*AuthHandler
	*BankHandler
}

func NewHandlers(
	authHandler *AuthHandler,
	bankHandler *BankHandler,
) *Handlers {
	return &Handlers{
		authHandler,
		bankHandler,
	}
}
