package entities

type ShoppingCart struct {
	// gorm.Model
	//SHOPPINGCARTID AUTO GENERATE
	ID        uint     `gorm:"primary_key:auto_increment" json:"id" form:"id"`
	User      User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    uint     `gorm:"not null" json:"user_id" form:"id"`
	ProductID uint     `gorm:"not null" json:"product_id" form:"product_id"`
	Order     []*Order `gorm:"many2many:shoppingcart_order;"`
	Qty       int      `json:"quantity" form:"quantity"`
}
