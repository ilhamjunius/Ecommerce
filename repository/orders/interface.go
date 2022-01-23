package order

import "ecommerce/entities"

type OrderInterface interface {
	GetAll(userid int) ([]entities.Order, error)
	Get(orderId, userId int) (entities.Order, error)
	Create(newOrder entities.Order, arr []int) (entities.Order, error)
	Cancel(orderId int, userId int) (entities.Order, error)
	Pay(invoiceId, paymentLink string, orderId, userId int) (entities.Order, error)
	Check(orderId int, userId int) (entities.Order, error)
}
