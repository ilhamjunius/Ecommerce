package entities

type Product struct {
	ID           uint `gorm:"productid;primary_key:auto_increment" json:"id" form:"id"`
	ShoppingCart []ShoppingCart
	Name         string `json:"product_name" form:"product_name"`
	Price        int    `json:"price" form:"price"`
	Stock        int    `json:"stock" form:"stok"`
	CategoryID   uint   `gorm:"not null" json:"category_id" form:"category_id"`
	Description  string `json:"description" form:"description"`
}
