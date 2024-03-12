package response

import "github.com/derangga/shopifyx/internal/entity"

type Auth struct {
	Name        string `json:"name"`
	Username    string `json:"username"`
	AccessToken string `json:"accessToken"`
}

func UserEntityToAuthResponse(data *entity.User) *Auth {
	return &Auth{
		Name:        data.Name,
		Username:    data.Username,
		AccessToken: data.AccessToken,
	}
}
