package response

import (
	"strings"

	"github.com/derangga/shopifyx/internal/entity"
)

type Product struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	Price          int      `json:"price"`
	ImageURL       string   `json:"imageUrl"`
	Stock          int      `json:"stock"`
	Condition      string   `json:"condition"`
	Tags           []string `json:"tags"`
	IsPurchaseable bool     `json:"isPurchaseable"`
}

func ProductEntityToResponse(data *entity.Product) *Product {
	return &Product{
		ID:             data.ID,
		Name:           data.Name,
		Price:          data.Price,
		ImageURL:       data.ImageURL,
		Stock:          data.Stock,
		Condition:      strings.ToLower(data.Condition),
		Tags:           data.Tags,
		IsPurchaseable: data.IsPurchaseable,
	}
}
