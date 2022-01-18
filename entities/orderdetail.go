package entities

import "gorm.io/gorm"

type OrdedDetail struct {
	gorm.Model
	ShoppingCartID uint
	OrderID        uint
}
