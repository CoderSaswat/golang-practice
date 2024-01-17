package service

import (
	"demo/model"
	"fmt"
)

type ProductServiceImpl struct {
	// Any additional fields related to the implementation can be added here
	// For simplicity, there are none in this example
}

// GetProductByID retrieves a product by its ID
func (p ProductServiceImpl) GetProductByID(id int) (*model.Product, error) {
	// In a real implementation, you would fetch the product from a data source
	// For simplicity, let's create a dummy product here
	return &model.Product{ID: id, Name: "Sample Product", Price: 19.99}, nil
}

func (p ProductServiceImpl) CreateProduct(product *model.Product) error {
	// In a real implementation, you would persist the product to a data source
	// For simplicity, let's just print the details in this example
	fmt.Printf("Product created: %+v\n", *product)
	return nil
}
