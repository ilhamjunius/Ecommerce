package entities

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	ID            uint   `json:"id" form:"id"`
	PaymentDetail string `json:"payment_detail" form:"payment_detail"`
	Status        string `json:"status" form:"status"`
	Order         Order  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
