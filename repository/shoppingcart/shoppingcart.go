package shoppingcart

import (
	"ecommerce/entities"
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ShoppingCartRepository struct {
	db *gorm.DB
}

func NewShoppingCartRepo(db *gorm.DB) *ShoppingCartRepository {
	return &ShoppingCartRepository{db: db}
}
func (ur *ShoppingCartRepository) Get(userId int) ([]entities.ShoppingCart, error) {
	cart := []entities.ShoppingCart{}
	if err := ur.db.Where("user_id = ?", userId).Find(&cart).Error; err != nil {
		log.Warn("Found database error", err)
		return cart, err
	}
	return cart, nil
}
func (ur *ShoppingCartRepository) Create(newShoppingcart entities.ShoppingCart) (entities.ShoppingCart, error) {
	product := entities.Product{}
	// type result struct {
	// 	Price int
	// }
	// res := ur.db.Joins(&).Find(&newShoppingcart)
	// ur.db.Model(&newShoppingcart).Select("products.price").Joins("left join products on shopping_carts.product_id = products.id").Scan(&product)
	price := ur.db.Find(&product, "id= ?", newShoppingcart.ProductID)
	fmt.Println("price:", price)
	newShoppingcart.Subtotal = newShoppingcart.Qty * product.Price

	// total:=
	// if err := ur.db.Model(&newShoppingcart).Preload("Product").Find(&newShoppingcart).Error; err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%+v\n", shoppingCart)
	// fmt.Println(shoppingCart.Product.Price)
	// newShoppingcart.Subtotal = uint(shoppingCart.Product.Price) * uint(newShoppingcart.Qty)

	if err := ur.db.Save(&newShoppingcart).Error; err != nil {
		return newShoppingcart, err
	}
	return newShoppingcart, nil
}
func (ur *ShoppingCartRepository) Update(updateCart entities.ShoppingCart, cartId int) (entities.ShoppingCart, error) {
	cart := entities.ShoppingCart{}
	if err := ur.db.First(&cart, "id=?", cartId).Error; err != nil {
		return cart, err
	}

	// cart.Qty = updateCart.Qty
	// cart.Subtotal = updateCart.Subtotal
	ur.db.Model(&cart).Updates(updateCart)

	return cart, nil
}
func (ur *ShoppingCartRepository) Delete(cartId int) (entities.ShoppingCart, error) {
	cart := entities.ShoppingCart{}
	if err := ur.db.Find(&cart, "id=?", cartId).Error; err != nil {
		return cart, err
	}
	if err := ur.db.Delete(&cart).Error; err != nil {
		return cart, err
	}
	return cart, nil
}
