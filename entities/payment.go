package entities

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	PaymentID     uint
	PaymentDetail string
	GrandTotal    int
	Status        string
}
