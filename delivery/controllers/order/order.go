package order

import order "ecommerce/repository/orders"

type OrderController struct {
	Repo order.OrderInterface
}

func NewOrderControllers(oi order.OrderInterface) *OrderController {
	return &OrderController{Repo: oi}
}
