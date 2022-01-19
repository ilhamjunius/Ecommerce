package shoppingcart

type ShoppingCartRequestFormat struct {
	UserId    uint `json:"user_id" form:"user_id"`
	ProductId uint `json:"product_id" form:"product_id"`
	Qty       int  `json:"quantity" form:"quantity"`
}
