package entities

type ShoppingCart struct {
	// gorm.Model
	//SHOPPINGCARTID AUTO GENERATE
	ID        uint `gorm:"primary_key:auto_increment" json:"shoppingcart_id" form:"shoppingcart_id"`
	OrderId   uint `json:"order_id" form:"order_id"`
	UserID    uint `gorm:"not null" json:"user_id" form:"user_id"`
	ProductID uint `gorm:"not null" json:"product_id" form:"product_id"`
	Qty       int  `json:"qty" form:"qty"`
	Subtotal  int  `json:"subtotal" form:"subtotal"`
}
