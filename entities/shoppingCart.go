package entities

import "gorm.io/gorm"

type ShoppingCart struct {
	gorm.Model
	//SHOPPINGCARTID AUTO GENERATE
	ShoppingCartID uint
	User           User
	Product        Product
	Qty            int
}
