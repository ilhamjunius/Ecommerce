package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID        uint
	PaymentId uint
	UserID    uint
	Total     int
	Status    string
}
