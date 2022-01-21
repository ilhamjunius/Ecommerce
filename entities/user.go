package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Email           string `gorm:"index:,unique" json:"email" form:"email"`
	Password        []byte `gorm:"not null" json:"password" form:"password"`
	Name            string `json:"name" form:"name"`
	HandphoneNumber string `json:"no_hp" form:"no_hp"`
	Role            string `json:"role" form:"role"`
	Order           []Order
	ShoppingCart    []ShoppingCart
}
