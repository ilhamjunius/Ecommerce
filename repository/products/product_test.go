package product

import (
	"ecommerce/configs"
	"ecommerce/entities"
	"ecommerce/repository/category"
	"ecommerce/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductRepo(t *testing.T) {
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

	categoryRepo := category.NewCategoryRepo(db)
	productRepo := NewProductRepo(db)

	var mockcategory entities.Category
	mockcategory.CategoryType = "mainan"
	categoryRepo.Create(mockcategory)

	var pagination entities.Pagination
	pagination.Limit = 1
	pagination.Page = 0
	pagination.Rows = 1

	t.Run("Insert Product into Database", func(t *testing.T) {
		var mockProduct entities.Product
		mockProduct.Name = "bola"
		mockProduct.Price = 10000
		mockProduct.Stock = 10
		mockProduct.CategoryID = 1
		mockProduct.Description = "bola basket"

		res, err := productRepo.Create(mockProduct)
		assert.Nil(t, err)
		assert.Equal(t, mockProduct.Name, res.Name)
		assert.Equal(t, 1, int(res.ID))

	})

	t.Run("Select Product from Database", func(t *testing.T) {
		res, err := productRepo.Get(1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})

	t.Run("Test Update Product ", func(t *testing.T) {
		var mockProduct entities.Product
		mockProduct.Name = "bola ajaib"
		mockProduct.Price = 50000
		mockProduct.Stock = 1
		mockProduct.CategoryID = 1
		mockProduct.Description = "bola sihir"

		res, err := productRepo.Update(mockProduct, 1)
		assert.Nil(t, err)
		assert.Equal(t, mockProduct.Name, res.Name)
		assert.Equal(t, 0, int(res.ID))
	})
	t.Run("Test Pagination", func(t *testing.T) {
		res, _ := productRepo.Pagination("bola", "mainan", pagination)
		assert.Nil(t, res.Error)
	})
	t.Run("Test Delete Product", func(t *testing.T) {
		_, err := productRepo.Delete(1)
		assert.Nil(t, err)
	})

	//DROP TABLE TESTING ERROR
	db.Migrator().DropTable(&entities.ShoppingCart{})
	db.Migrator().DropTable(&entities.Order{})
	db.Migrator().DropTable(&entities.Product{})
	db.Migrator().DropTable(&entities.Category{})
	db.Migrator().DropTable(&entities.User{})

	t.Run("Error Insert Product into Database", func(t *testing.T) {
		var mockProduct entities.Product
		mockProduct.Name = "bola"
		mockProduct.Price = 10000
		mockProduct.Stock = 10
		mockProduct.CategoryID = 2
		mockProduct.Description = "bola basket"
		_, err := productRepo.Create(mockProduct)
		assert.NotNil(t, err)
	})

	t.Run("Error Select Product from Database", func(t *testing.T) {
		_, err := productRepo.Get(1)
		assert.NotNil(t, err)
	})

	t.Run("Update User ", func(t *testing.T) {
		var mockProduct entities.Product
		mockProduct.Name = "bola ajaib"
		mockProduct.Price = 50000
		mockProduct.Stock = 1
		mockProduct.CategoryID = 1
		mockProduct.Description = "bola sihir"

		_, err := productRepo.Update(mockProduct, 1)
		assert.NotNil(t, err)
	})

	t.Run("Error Delete User", func(t *testing.T) {
		_, err := productRepo.Delete(1)
		assert.NotNil(t, err)
	})
}
