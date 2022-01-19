package shoppingcart

import "ecommerce/entities"

type ProductResponseFormat struct {
	Message string                `json:"message"`
	Data    entities.ShoppingCart `json:"data"`
}
