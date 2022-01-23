package shoppingcart

import "ecommerce/entities"

type ShoppingCartInterface interface {
	Get(userId int) ([]entities.ShoppingCart, error)
	GetById(Id, userId int) (entities.ShoppingCart, error)
	Create(newCart entities.ShoppingCart) (entities.ShoppingCart, error)
	Update(updateCart entities.ShoppingCart, cartId, userId int) (entities.ShoppingCart, error)
	Delete(cartId, userId int) (entities.ShoppingCart, error)
}
