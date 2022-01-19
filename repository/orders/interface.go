package order

import "ecommerce/entities"

type OrderInterface interface {
	GetAll(userid int) ([]entities.Order, error)
	Get(orderId, userId int) (entities.Order, error)
	Create(newOrder entities.Order) (entities.Order, error)
	Update(newOrder entities.Order, orderId, userId int) (entities.Order, error)
	Delete(orderId, userId int) (entities.Order, error)
}
