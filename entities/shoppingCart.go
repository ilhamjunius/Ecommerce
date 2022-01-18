package entities

import "gorm.io/gorm"

type ShoppingCart struct {
	gorm.Model
	//SHOPPINGCARTID AUTO GENERATE
	ID        uint
	User      User
	UserID    uint
	Product   Product
	ProductID uint
	Order     []*Order `gorm:"many2many:shoppingcart_order;"`
	Qty       int
}
