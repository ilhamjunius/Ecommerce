package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	//PRODUCTID AUTO GENERATE
	ID          uint     `json:"id" form:"id"`
	Name        string   `json:"name" form:"name"`
	Price       int      `json:"price" form:"price"`
	Stock       int      `json:"stok" form:"stok"`
	CategoryID  uint     `gorm:"unique;not null" json:"category_id" form:"category_id"`
	Category    Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Description string   `json:"description" form:"description"`
}
