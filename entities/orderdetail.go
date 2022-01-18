package entities

import "gorm.io/gorm"

type OrdedDetail struct {
	gorm.Model
	Order        Order
	ShoppingCart []ShoppingCart
}
