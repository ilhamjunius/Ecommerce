package entities

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID           uint      `gorm:"primary_key:auto_increment" json:"id" form:"id"`
	CategoryType string    `gorm:"unique;not null" json:"category_type" form:"category_type"`
	Product      []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
