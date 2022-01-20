package shoppingcart

import (
	"ecommerce/configs"
	"ecommerce/entities"
	"ecommerce/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShoppingCartRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.ShoppingCart{})
	db.AutoMigrate(&entities.ShoppingCart{})

	shoppingCartRepo := NewShoppingCartRepo(db)
	t.Run("Insert ShoppingCart into Database", func(t *testing.T) {
		var mockShoppingCart entities.ShoppingCart
		mockShoppingCart.ID = 1
		mockShoppingCart.OrderId = 1
		mockShoppingCart.UserID = 2
		mockShoppingCart.ProductID = 1
		mockShoppingCart.Qty = 2
		mockShoppingCart.Subtotal = 10000
		db.Migrator().DropTable(&entities.ShoppingCart{})
		_, err := shoppingCartRepo.Create(mockShoppingCart)
		assert.Error(t, err)

	})
	db.AutoMigrate(&entities.ShoppingCart{})
	t.Run("Insert ShoppingCart into Database", func(t *testing.T) {
		var mockShoppingCart entities.ShoppingCart
		mockShoppingCart.OrderId = 1
		mockShoppingCart.UserID = 2
		mockShoppingCart.ProductID = 1
		mockShoppingCart.Qty = 2
		mockShoppingCart.Subtotal = 10000

		res, err := shoppingCartRepo.Create(mockShoppingCart)
		assert.Nil(t, err)
		assert.Equal(t, mockShoppingCart.Qty, res.Qty)
		assert.Equal(t, 1, int(res.ID))

	})
	t.Run("Select ShoppingCart from Database", func(t *testing.T) {

		res, err := shoppingCartRepo.Get(1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})
	t.Run("Select ShoppingCart from Database", func(t *testing.T) {
		db.Migrator().DropTable(&entities.ShoppingCart{})
		_, err := shoppingCartRepo.Get(1)
		assert.Error(t, err)

	})
	db.AutoMigrate(&entities.ShoppingCart{})
	t.Run("Insert ShoppingCart into Database", func(t *testing.T) {
		var mockShoppingCart entities.ShoppingCart
		mockShoppingCart.OrderId = 1
		mockShoppingCart.UserID = 2
		mockShoppingCart.ProductID = 1
		mockShoppingCart.Qty = 2
		mockShoppingCart.Subtotal = 10000

		res, err := shoppingCartRepo.Create(mockShoppingCart)
		assert.Nil(t, err)
		assert.Equal(t, mockShoppingCart.Qty, res.Qty)
		assert.Equal(t, 1, int(res.ID))

	})
	t.Run("Update ShoppingCart ", func(t *testing.T) {
		var mockCategory entities.ShoppingCart
		mockCategory.Qty = 4
		res, err := shoppingCartRepo.Update(mockCategory, 1)
		assert.Nil(t, err)
		assert.Equal(t, mockCategory.Qty, res.Qty)
	})
	t.Run("Error Update ShoppingCart ", func(t *testing.T) {
		var mockCategory entities.ShoppingCart
		mockCategory.Qty = 4
		db.Migrator().DropTable(&entities.ShoppingCart{})
		_, err := shoppingCartRepo.Update(mockCategory, 1)
		assert.Error(t, err)

	})
	t.Run("Error Delete ShoppingCart", func(t *testing.T) {
		_, err := shoppingCartRepo.Delete(1)
		assert.Error(t, err)

	})
	db.AutoMigrate(&entities.ShoppingCart{})
	t.Run("Insert ShoppingCart into Database", func(t *testing.T) {
		var mockShoppingCart entities.ShoppingCart
		mockShoppingCart.OrderId = 1
		mockShoppingCart.UserID = 2
		mockShoppingCart.ProductID = 1
		mockShoppingCart.Qty = 2
		mockShoppingCart.Subtotal = 10000

		res, err := shoppingCartRepo.Create(mockShoppingCart)
		assert.Nil(t, err)
		assert.Equal(t, mockShoppingCart.Qty, res.Qty)
		assert.Equal(t, 1, int(res.ID))

	})
	t.Run("Delete ShoppingCart", func(t *testing.T) {
		res, err := shoppingCartRepo.Delete(1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})

}
