package product

import "ecommerce/entities"

type ProductInterface interface {
	// GetAll(keyword string) ([]entities.Product, error)
	Pagination(name, category string, pagination entities.Pagination) (ResultPagination, int)
	Get(productId int) (entities.Product, error)
	Create(newProduct entities.Product) (entities.Product, error)
	Delete(productId int) (entities.Product, error)
	Update(newProduct entities.Product, productId int) (entities.Product, error)
}

type ResultPagination struct {
	Result entities.Pagination
	Error  error
}
