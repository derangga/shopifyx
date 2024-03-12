package handler

type Handlers struct {
	*AuthHandler
}

func NewHandlers(
	authHandler *AuthHandler,
) *Handlers {
	return &Handlers{
		authHandler,
	}
}
