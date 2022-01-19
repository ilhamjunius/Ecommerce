package shoppingcart

import "ecommerce/entities"

type ShoppingCartInterface interface {
	Get(userIsd int) (entities.ShoppingCart, error)
	Create(newCart entities.ShoppingCart) (entities.ShoppingCart, error)
	Update(quantity, cartId int) (entities.ShoppingCart, error)
	Delete(cartId int) (entities.ShoppingCart, error)
}
