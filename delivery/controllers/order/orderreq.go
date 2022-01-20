package order

type OrderCancelRequestFormat struct {
	PaymentID uint `json:"payment_id" form:"payment_id"`
	Qty       int  `json:"quantity" form:"quantity"`
}
