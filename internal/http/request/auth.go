package request

import "github.com/derangga/shopifyx/internal/entity"

type Register struct {
	Username string `json:"username" validate:"required,min=5,max=15"`
	Name     string `json:"name"     validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=5,max=15"`
}

func (r *Register) ToEntityUser() *entity.User {
	return &entity.User{
		Username: r.Username,
		Name:     r.Name,
		Password: r.Password,
	}
}

type Login struct {
	Username string `json:"username" validate:"required,min=5,max=15"`
	Password string `json:"password" validate:"required,min=5,max=15"`
}

func (r *Login) ToEntityUser() *entity.User {
	return &entity.User{
		Username: r.Username,
		Password: r.Password,
	}
}
