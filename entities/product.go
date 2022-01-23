package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID           uint   `json:"product_id" form:"product_id"`
	Name         string `json:"product_name" form:"product_name"`
	Price        int    `json:"price" form:"price"`
	Stock        int    `json:"stock" form:"stok"`
	Description  string `json:"description" form:"description"`
	CategoryID   uint   `gorm:"not null" json:"category_id" form:"category_id"`
	ShoppingCart []ShoppingCart
}
