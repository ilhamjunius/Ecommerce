package order

import "ecommerce/entities"

type GetOrdersResponseFormat struct {
	Message string           `json:"message"`
	Data    []entities.Order `json:"data"`
}

type GetOrderResponseFormat struct {
	Message string         `json:"message"`
	Data    entities.Order `json:"data"`
}

type CreateOrderResponseFormat struct {
	Message string         `json:"message"`
	Data    entities.Order `json:"data"`
}

type CancelOrderResponseFormat struct {
	Message string         `json:"message"`
	Data    entities.Order `json:"data"`
}
