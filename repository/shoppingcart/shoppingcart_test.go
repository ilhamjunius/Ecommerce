package shoppingcart

import (
	"crypto/sha256"
	"ecommerce/configs"
	"ecommerce/entities"
	"ecommerce/repository/category"
	product "ecommerce/repository/products"
	user "ecommerce/repository/users"
	"ecommerce/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShoppingCartRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.ShoppingCart{})
	db.Migrator().DropTable(&entities.Order{})
	db.Migrator().DropTable(&entities.Product{})
	db.Migrator().DropTable(&entities.Category{})
	db.Migrator().DropTable(&entities.User{})

	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Order{})
	db.AutoMigrate(&entities.Category{})
	db.AutoMigrate(&entities.Product{})
	db.AutoMigrate(&entities.ShoppingCart{})

	userRepo := user.NewUserRepo(db)
	categoryRepo := category.NewCategoryRepo(db)
	productRepo := product.NewProductRepo(db)
	shoppingCartRepo := NewShoppingCartRepo(db)

	var mockUser entities.User
	hash := sha256.Sum256([]byte("ilham123"))
	mockUser.Email = "ilham@gmail.com"
	mockUser.Password = hash[:]
	mockUser.HandphoneNumber = "0123456789"
	mockUser.Role = "admin"
	userRepo.Create(mockUser)

	var mockcategory entities.Category
	mockcategory.CategoryType = "mainan"
	categoryRepo.Create(mockcategory)

	var mockproduct entities.Product
	mockproduct.Name = "bola ajaib"
	mockproduct.Price = 5000
	mockproduct.Stock = 50
	mockproduct.CategoryID = 1
	mockproduct.Description = "sangat ajaib"
	productRepo.Create(mockproduct)

	t.Run("Insert ShoppingCart into Database", func(t *testing.T) {
		var mockShoppingCart entities.ShoppingCart
		mockShoppingCart.UserID = 1
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
	t.Run("Update ShoppingCart ", func(t *testing.T) {
		var mockCategory entities.ShoppingCart
		mockCategory.Qty = 4
		res, err := shoppingCartRepo.Update(mockCategory, 1, 1)
		assert.Nil(t, err)
		assert.Equal(t, mockCategory.Qty, res.Qty)
	})
	t.Run("Delete ShoppingCart", func(t *testing.T) {
		res, err := shoppingCartRepo.Delete(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})

	//DROP TABLE UNTUK ERROR
	db.Migrator().DropTable(&entities.ShoppingCart{})
	db.Migrator().DropTable(&entities.Order{})
	db.Migrator().DropTable(&entities.Product{})
	db.Migrator().DropTable(&entities.Category{})
	db.Migrator().DropTable(&entities.User{})

	t.Run("Insert ShoppingCart into Database", func(t *testing.T) {
		var mockShoppingCart entities.ShoppingCart
		mockShoppingCart.UserID = 1
		mockShoppingCart.ProductID = 1
		mockShoppingCart.Qty = 2
		mockShoppingCart.Subtotal = 10000

		_, err := shoppingCartRepo.Create(mockShoppingCart)
		assert.NotNil(t, err)
	})
	t.Run("Error Select ShoppingCart from Database", func(t *testing.T) {
		_, err := shoppingCartRepo.Get(1)
		assert.NotNil(t, err)
	})
	t.Run("Error Update ShoppingCart ", func(t *testing.T) {
		var mockCategory entities.ShoppingCart
		mockCategory.Qty = 4
		_, err := shoppingCartRepo.Update(mockCategory, 1, 1)
		assert.NotNil(t, err)
	})
	t.Run("Error Delete ShoppingCart", func(t *testing.T) {
		_, err := shoppingCartRepo.Delete(1, 1)
		assert.NotNil(t, err)
	})
}
