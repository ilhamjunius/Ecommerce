package product

import "ecommerce/entities"

type ProductInterface interface {
	GetAll() ([]entities.Product, error)
	Get(productId int) (entities.Product, error)
	Create(newProduct entities.Product) (entities.Product, error)
	Delete(productId int) (entities.Product, error)
	Update(newProduct entities.Product, productId int) (entities.Product, error)
}
