package contract

import "demo/model"

// ProductService is an interface for interacting with products
type ProductService interface {
	GetProductByID(id int) (*model.Product, error)
	CreateProduct(product *model.Product) error
}
