package entities

type OrdedDetail struct {
	// gorm.Model
	Order        Order
	ShoppingCart []ShoppingCart
}
