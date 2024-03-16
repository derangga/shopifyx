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

type ProductDetail struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	Price          int      `json:"price"`
	ImageURL       string   `json:"imageUrl"`
	Stock          int      `json:"stock"`
	Condition      string   `json:"condition"`
	Tags           []string `json:"tags"`
	IsPurchaseable bool     `json:"isPurchaseable"`
	PurchaseCount  int      `json:"purchaseCount"`
	Seller         Seller
}

type Seller struct {
	Name             string `json:"name"`
	ProductSoldTotal int    `json:"productSoldTotal"`
	BankAccount      SellerBank
}

type SellerBank struct {
	ID                int    `json:"id"`
	BankName          string `json:"bankName"`
	BankAccountName   string `json:"bankAccountName"`
	BankAccountNumber string `json:"bankAccountNumber"`
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

func ProductDetailToResponse(data *entity.ProductDetail) *ProductDetail {
	bankAccount := SellerBank{
		ID:                data.Seller.BankAccount.ID,
		BankName:          data.Seller.BankAccount.BankName,
		BankAccountName:   data.Seller.BankAccount.BankAccountName,
		BankAccountNumber: data.Seller.BankAccount.BankAccountNumber,
	}
	seller := Seller{
		Name:             data.Seller.Name,
		ProductSoldTotal: data.Seller.ProductSoldTotal,
		BankAccount:      bankAccount,
	}
	return &ProductDetail{
		ID:             data.ID,
		Name:           data.Name,
		Price:          data.Price,
		ImageURL:       data.ImageURL,
		Stock:          data.Stock,
		Condition:      data.Condition,
		Tags:           data.Tags,
		IsPurchaseable: data.IsPurchaseable,
		PurchaseCount:  data.PurchaseCount,
		Seller:         seller,
	}
}
