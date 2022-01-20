package product

import "ecommerce/entities"

type ProductInterface interface {
	GetAll(keyword string) ([]entities.Product, error)
	Get(productId int) (entities.Product, error)
	FilterProduct(keyword, category string) ([]entities.Product, error)
	Create(newProduct entities.Product) (entities.Product, error)
	Delete(productId int) (entities.Product, error)
	Update(newProduct entities.Product, productId int) (entities.Product, error)
}
