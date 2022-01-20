package shoppingcart

type ShoppingCartRequestFormat struct {
	OrderId   uint `json:"order_id" form:"order_id"`
	UserId    uint `json:"user_id" form:"user_id"`
	ProductId uint `json:"product_id" form:"product_id"`
	Qty       int  `json:"quantity" form:"quantity"`
	Subtotal  int  `json:"subtotal" form:"subtotal"`
}
