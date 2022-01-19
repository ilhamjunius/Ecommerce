package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID        uint   `gorm:"primary_key:auto_increment" json:"id" form:"id"`
	PaymentId uint   `gorm:"unique;not null" json:"payment_id" form:"payment_id"`
	UserID    uint   `gorm:"unique;not null" json:"user_id" form:"user_id"`
	Total     int    `json:"total" form:"total"`
	Status    string `json:"status" form:"status"`
}
