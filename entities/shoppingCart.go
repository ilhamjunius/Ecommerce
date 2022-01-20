package entities

import (
	"time"

	"gorm.io/gorm"
)

type ShoppingCart struct {
	ID        uint `gorm:"primary_key:auto_increment" json:"shoppingcart_id" form:"shoppingcart_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	OrderId   uint           `gorm:"default:NULL" json:"order_id" form:"order_id"`
	UserID    uint           `gorm:"not null" json:"user_id" form:"user_id"`
	ProductID uint           `gorm:"not null" json:"product_id" form:"product_id"`
	Qty       int            `json:"qty" form:"qty"`
	Subtotal  uint           `json:"subtotal" form:"subtotal"`
}
