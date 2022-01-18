package entities

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID           uint   `json:"id" form:"id"`
	CategoryType string `gorm:"unique;not null" json:"category_type" form:"category_type"`
}
