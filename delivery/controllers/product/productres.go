package product

import "ecommerce/entities"

type ProductResponseFormat struct {
	Message string           `json:"message"`
	Data    entities.Product `json:"data"`
}

type GetAllProductsResponseFormat struct {
	Message string             `json:"message"`
	Data    []entities.Product `json:"data"`
}
type PaginationResponseFormat struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   error
}
