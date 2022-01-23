package entities

import (
	"gorm.io/gorm"
)

type ShoppingCart struct {
	gorm.Model
	ID        uint `json:"shoppingcart_id" form:"shoppingcart_id"`
	OrderId   uint `gorm:"default:NULL" json:"order_id" form:"order_id"`
	ProductID uint `gorm:"not null" json:"product_id" form:"product_id"`
	UserID    uint `gorm:"not null" json:"user_id" form:"user_id"`
	Qty       int  `json:"quantity" form:"quantity"`
	Subtotal  int  `json:"subtotal" form:"subtotal"`
}
