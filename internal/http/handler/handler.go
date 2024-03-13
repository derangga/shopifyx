package handler

type Handlers struct {
	*AuthHandler
	*ProductHandler
}

func NewHandlers(
	authHandler *AuthHandler,
	productHandler *ProductHandler,
) *Handlers {
	return &Handlers{
		authHandler,
		productHandler,
	}
}
