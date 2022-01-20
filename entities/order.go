package entities

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	// gorm.Model
	ID           uint `gorm:"primary_key:auto_increment" json:"order_id" form:"order_id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	ShoppingCart []ShoppingCart
	PaymentId    uint   `gorm:"default:NULL" json:"payment_id" form:"payment_id"`
	UserID       uint   `gorm:"unique;not null" json:"user_id" form:"user_id"`
	Total        int    `json:"total" form:"total"`
	Status       string `json:"status" form:"status"`
}
