package entities

import "gorm.io/gorm"

type ShoppingCart struct {
	gorm.Model
	//SHOPPINGCARTID AUTO GENERATE
	ShoppingCartID uint
	UserID         uint
	ProductID      uint
	Qty            int
}
