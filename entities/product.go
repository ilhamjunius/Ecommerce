package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	//PRODUCTID AUTO GENERATE
	ProductID   uint
	Name        string
	Price       int
	Stock       int
	Category    Category
	Description string
}
