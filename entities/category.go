package entities

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryID   uint
	CategoryType string
}
