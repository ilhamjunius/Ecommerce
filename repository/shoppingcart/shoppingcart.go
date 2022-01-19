package shoppingcart

import (
	"ecommerce/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ShoppingCartRepository struct {
	db *gorm.DB
}

func NewShoppingCartRepo(db *gorm.DB) *ShoppingCartRepository {
	return &ShoppingCartRepository{db: db}
}
func (ur *ShoppingCartRepository) Get(userId int) (entities.ShoppingCart, error) {
	cart := entities.ShoppingCart{}
	if err := ur.db.Find(&cart, userId).Error; err != nil {
		log.Warn("Found database error", err)
		return cart, err
	}
	return cart, nil
}
func (ur *ShoppingCartRepository) Create(newShoppingcart entities.ShoppingCart) (entities.ShoppingCart, error) {
	if err := ur.db.Save(&newShoppingcart).Error; err != nil {
		return newShoppingcart, err
	}
	return newShoppingcart, nil
}
func (ur *ShoppingCartRepository) Update(quantity, cartId int) (entities.ShoppingCart, error) {
	cart := entities.ShoppingCart{}
	if err := ur.db.Find(&cart, "id=?", cartId).Error; err != nil {
		return cart, err
	}

	cart.Qty = quantity

	if err := ur.db.Save(&cart).Error; err != nil {
		return cart, err
	}

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
