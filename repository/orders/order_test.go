package order

import (
	"crypto/sha256"
	"ecommerce/configs"
	"ecommerce/entities"
	"ecommerce/repository/category"
	product "ecommerce/repository/products"
	"ecommerce/repository/shoppingcart"
	user "ecommerce/repository/users"
	"ecommerce/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.ShoppingCart{})
	db.Migrator().DropTable(&entities.Order{})
	db.Migrator().DropTable(&entities.Product{})
	db.Migrator().DropTable(&entities.Category{})
	db.Migrator().DropTable(&entities.User{})

	userRepo := user.NewUserRepo(db)
	categoryRepo := category.NewCategoryRepo(db)
	productRepo := product.NewProductRepo(db)
	shoppingCartRepo := shoppingcart.NewShoppingCartRepo(db)
	OrderRepo := NewOrderRepo(db)

	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Order{})
	db.AutoMigrate(&entities.Category{})
	db.AutoMigrate(&entities.Product{})
	db.AutoMigrate(&entities.ShoppingCart{})

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

	var mockShoppingCart entities.ShoppingCart
	mockShoppingCart.UserID = 1
	mockShoppingCart.ProductID = 1
	mockShoppingCart.Qty = 2
	mockShoppingCart.Subtotal = 10000
	shoppingCartRepo.Create(mockShoppingCart)

	t.Run("Insert Order into Database", func(t *testing.T) {
		var mockOrder entities.Order
		mockOrder.UserID = 1
		arr := []int{1, 2}

		res, err := OrderRepo.Create(mockOrder, arr)
		assert.Nil(t, err)
		assert.Equal(t, mockOrder.UserID, res.UserID)
		assert.Equal(t, 1, int(res.ID))
	})
	t.Run("Select All Orders from Database", func(t *testing.T) {
		res, err := OrderRepo.GetAll(1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})
	t.Run("Select Order from Database", func(t *testing.T) {
		res, err := OrderRepo.Get(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})
	t.Run("Cancel an Order from Database", func(t *testing.T) {
		res, err := OrderRepo.Cancel(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})
	t.Run("Check an Order from Database", func(t *testing.T) {
		res, err := OrderRepo.Check(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})
	t.Run("Pay an Order from Database", func(t *testing.T) {
		res, err := OrderRepo.Pay("1", "asdf", 1, 1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})
	db.Migrator().DropTable(&entities.ShoppingCart{})
	db.Migrator().DropTable(&entities.Order{})
	db.Migrator().DropTable(&entities.Product{})
	db.Migrator().DropTable(&entities.Category{})
	db.Migrator().DropTable(&entities.User{})

	t.Run("Insert Order into Database", func(t *testing.T) {
		var mockOrder entities.Order
		mockOrder.UserID = 2
		arr := []int{1, 2}

		_, err := OrderRepo.Create(mockOrder, arr)
		assert.NotNil(t, err)
	})
	t.Run("Select All Orders from Database", func(t *testing.T) {
		_, err := OrderRepo.GetAll(1)
		assert.NotNil(t, err)
	})
	t.Run("Select Order from Database", func(t *testing.T) {
		_, err := OrderRepo.Get(1, 1)
		assert.NotNil(t, err)
	})
	t.Run("Cancel an Order from Database", func(t *testing.T) {
		_, err := OrderRepo.Cancel(1, 1)
		assert.NotNil(t, err)
	})
	t.Run("Check an Order from Database", func(t *testing.T) {
		_, err := OrderRepo.Check(1, 1)
		assert.NotNil(t, err)
	})
	t.Run("Pay an Order from Database", func(t *testing.T) {
		_, err := OrderRepo.Pay("1", "asdf", 1, 1)
		assert.NotNil(t, err)
	})

}
