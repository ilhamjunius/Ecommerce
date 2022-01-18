package entities

import "gorm.io/gorm"

type ShoppingCart struct {
	gorm.Model
	//SHOPPINGCARTID AUTO GENERATE
	ID        uint `json:"id" form:"id"`
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    uint `gorm:"unique;not null" json:"user_id" form:"id"`
	Product   Product
	ProductID uint     `gorm:"unique;not null" json:"product_id" form:"product_id"`
	Order     []*Order `gorm:"many2many:shoppingcart_order;"`
	Qty       int      `json:"qty" form:"qty"`
}
