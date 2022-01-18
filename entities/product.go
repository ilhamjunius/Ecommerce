package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	//PRODUCTID AUTO GENERATE
	ID          uint
	Name        string
	Price       int
	Stock       int
	CategoryID  uint
	Category    Category
	Description string
}
