package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderID uint
	Payment Payment
	User    User
	Total   int
	Status  string
}
