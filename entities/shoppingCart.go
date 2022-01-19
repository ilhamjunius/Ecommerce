package entities

type ShoppingCart struct {
	// gorm.Model
	//SHOPPINGCARTID AUTO GENERATE
<<<<<<< HEAD
	ID        uint `gorm:"primary_key:auto_increment" json:"shoppingcart_id" form:"shoppingcart_id"`
	OrderId   uint `json:"order_id" form:"order_id"`
	UserID    uint `gorm:"not null" json:"user_id" form:"user_id"`
	ProductID uint `gorm:"not null" json:"product_id" form:"product_id"`
	Qty       int  `json:"qty" form:"qty"`
	Subtotal  uint `json:"subtotal" form:"subtotal"`
=======
	ID        uint     `gorm:"primary_key:auto_increment" json:"id" form:"id"`
	User      User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    uint     `gorm:"not null" json:"user_id" form:"id"`
	ProductID uint     `gorm:"not null" json:"product_id" form:"product_id"`
	Order     []*Order `gorm:"many2many:shoppingcart_order;"`
	Qty       int      `json:"quantity" form:"quantity"`
>>>>>>> aa8d0770fd5b2ea05a390922d7ce5824339d9f63
}
