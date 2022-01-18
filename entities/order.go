package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderID   uint
	PaymentID uint
	UserID    uint
	Total     int
	Status    string
}
