package entities

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	ID            uint
	PaymentDetail string
	Status        string
	Order         Order
}
