package product

type ProductRequestFormat struct {
	Name        string `json:"product_name" form:"product_name"`
	Price       int    `json:"price" form:"price"`
	Stock       int    `json:"stock" form:"stock"`
	CategoryID  uint   `json:"category_id" form:"category_id"`
	Description string `json:"description" form:"description"`
}
