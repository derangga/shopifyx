package request

import (
	"strings"

	"github.com/derangga/shopifyx/internal/entity"
)

type Product struct {
	Name           string   `json:"name"           validate:"required,min=5,max=50"`
	Price          int      `json:"price"          validate:"required,min=0"`
	ImageURL       string   `json:"imageUrl"       validate:"required,url"`
	Stock          int      `json:"stock"          validate:"required,min=0"`
	Condition      string   `json:"condition"      validate:"required,oneof=new second NEW SECOND"`
	Tags           []string `json:"tags"           validate:"required,min=0"`
	IsPurchaseable *bool    `json:"isPurchaseable" validate:"required"`
}

type ProductUpdate struct {
	ID             int      `param:"id" validate:"required"`
	Name           string   `json:"name"           validate:"required,min=5,max=15"`
	Price          int      `json:"price"          validate:"required,min=0"`
	ImageURL       string   `json:"imageUrl"       validate:"required,url"`
	Condition      string   `json:"condition"      validate:"required,oneof=new second NEW SECOND"`
	Tags           []string `json:"tags"           validate:"required,min=0"`
	IsPurchaseable *bool    `json:"isPurchaseable" validate:"required"`
}

type DeleteProduct struct {
	ID int `param:"id" validate:"required"`
}

type UpdateStock struct {
	ID    int `param:"id" validate:"required"`
	Stock int `json:"stock"          validate:"required,min=0"`
}

type ListFilter struct {
	Tags []string `query:"tags"`

	Condition string `query:"condition" validate:"oneof=new second"`
	Search    string `query:"search"`

	SortBy  string `query:"sortBy"`
	OrderBy string `query:"orderBy"`
	Page    int    `query:"offset" default:"1"`
	Limit   int    `query:"limit" default:"15"`

	MaxPrice int `query:"maxPrice"`
	MinPrice int `query:"minPrice"`

	UserOnly       bool `query:"userOnly" default:"false"`
	ShowEmptyStock bool `query:"showEmptyStock" default:"false"`
}

func (p *Product) ToEntityProduct() *entity.Product {
	return &entity.Product{
		Name:           p.Name,
		Price:          p.Price,
		ImageURL:       p.ImageURL,
		Stock:          p.Stock,
		Condition:      strings.ToUpper(p.Condition),
		Tags:           p.Tags,
		IsPurchaseable: *p.IsPurchaseable,
	}
}

func (p *ProductUpdate) ToEntityProduct() *entity.Product {
	return &entity.Product{
		ID:             p.ID,
		Name:           p.Name,
		Price:          p.Price,
		ImageURL:       p.ImageURL,
		Condition:      strings.ToUpper(p.Condition),
		Tags:           p.Tags,
		IsPurchaseable: *p.IsPurchaseable,
	}
}

func (p *DeleteProduct) ToEntityProduct() *entity.Product {
	return &entity.Product{
		ID: p.ID,
	}
}

func (p *UpdateStock) ToEntityProduct() *entity.Product {
	return &entity.Product{
		ID:    p.ID,
		Stock: p.Stock,
	}
}

func (p *ListFilter) ToEntityListFilter() *entity.ListFilter {
	return &entity.ListFilter{
		UserOnly:       p.UserOnly,
		Tags:           p.Tags,
		Condition:      strings.ToUpper(p.Condition),
		ShowEmptyStock: p.ShowEmptyStock,
		MaxPrice:       p.MaxPrice,
		MinPrice:       p.MinPrice,
		Search:         p.Search,
		SortBy:         p.SortBy,
		OrderBy:        p.OrderBy,
		Page:           p.Page,
		Limit:          p.Limit,
	}
}
