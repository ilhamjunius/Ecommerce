package entities

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	// gorm.Model
	//PRODUCTID AUTO GENERATE
	ID           uint `gorm:"productid;primary_key:auto_increment" json:"id" form:"id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Name         string         `json:"product_name" form:"product_name"`
	Price        int            `json:"price" form:"price"`
	Stock        int            `json:"stock" form:"stok"`
	CategoryID   uint           `gorm:"unique;not null" json:"category_id" form:"category_id"`
	Description  string         `json:"description" form:"description"`
	ShoppingCart []ShoppingCart `gorm:"-"`
}
