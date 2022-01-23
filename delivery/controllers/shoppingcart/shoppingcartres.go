package shoppingcart

import "ecommerce/entities"

type ShoppingCartResponseFormat struct {
	Message string                `json:"message"`
	Data    entities.ShoppingCart `json:"data"`
}
type ManyShoppingCartResponseFormat struct {
	Message string                  `json:"message"`
	Data    []entities.ShoppingCart `json:"data"`
}
