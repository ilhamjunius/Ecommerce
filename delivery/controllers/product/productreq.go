package product

type ProductRequestFormat struct {
	Name        string `json:"product_name" form:"product_name"`
	Price       string `json:"price" form:"price"`
	Stock       string `json:"stock" form:"stock"`
	CategoryID  string `json:"category_id" form:"category_id"`
	Description string `json:"description" form:"description"`
}
