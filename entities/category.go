package entities

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID           uint   `json:"category_id" form:"category_id"`
	CategoryType string `gorm:"unique;not null" json:"category_type" form:"category_type"`
	Product      []Product
	// `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
