package entities

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryType string    `gorm:"unique;not null" json:"category_type" form:"category_type"`
	Products     []Product `gorm:"foreignKey:CategoryID"`
}
