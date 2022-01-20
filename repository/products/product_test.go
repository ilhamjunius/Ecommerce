package product

import (
	"ecommerce/configs"
	"ecommerce/entities"
	"ecommerce/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Product{})
	db.AutoMigrate(&entities.Product{})

	productRepo := NewProductRepo(db)

	t.Run("Insert Product into Database", func(t *testing.T) {
		var mockProduct entities.Product
		mockProduct.Name = "bola"
		mockProduct.Price = 10000
		mockProduct.Stock = 10
		mockProduct.CategoryID = 2
		mockProduct.Description = "bola basket"

		res, err := productRepo.Create(mockProduct)
		assert.Nil(t, err)
		assert.Equal(t, mockProduct.Name, res.Name)
		assert.Equal(t, 1, int(res.ID))

	})

	t.Run("Select Products from Database", func(t *testing.T) {
		res, err := productRepo.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})

	t.Run("Error Insert Product into Database", func(t *testing.T) {
		var mockProduct entities.Product
		mockProduct.Name = "bola"
		mockProduct.Price = 10000
		mockProduct.Stock = 10
		mockProduct.CategoryID = 2
		mockProduct.Description = "bola basket"
		db.Migrator().DropTable(&entities.Product{})
		_, err := productRepo.Create(mockProduct)
		assert.Error(t, err)
	})

	t.Run("Error Select Products from Database", func(t *testing.T) {
		_, err := productRepo.GetAll()
		assert.Error(t, err)
	})

	t.Run("Error Select Product from Database", func(t *testing.T) {
		_, err := productRepo.Get(1)
		assert.Error(t, err)

	})
	t.Run("Error Update User ", func(t *testing.T) {
		var mockProduct entities.Product
		mockProduct.Name = "bola ajaib"
		mockProduct.Price = 50000
		mockProduct.Stock = 1
		mockProduct.CategoryID = 1
		mockProduct.Description = "bola sihir"

		_, err := productRepo.Update(mockProduct, 1)
		assert.Error(t, err)

	})
	t.Run("Error Delete User", func(t *testing.T) {
		_, err := productRepo.Delete(1)
		assert.Error(t, err)

	})

	t.Run("Insert Product into Database", func(t *testing.T) {
		db.AutoMigrate(&entities.Product{})
		var mockProduct entities.Product
		mockProduct.Name = "bola"
		mockProduct.Price = 10000
		mockProduct.Stock = 10
		mockProduct.CategoryID = 2
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

	t.Run("Update User ", func(t *testing.T) {
		var mockProduct entities.Product
		mockProduct.Name = "bola ajaib"
		mockProduct.Price = 50000
		mockProduct.Stock = 1
		mockProduct.CategoryID = 1
		mockProduct.Description = "bola sihir"

		res, err := productRepo.Update(mockProduct, 1)
		assert.Nil(t, err)
		assert.Equal(t, mockProduct.Name, res.Name)
	})

	t.Run("Delete User", func(t *testing.T) {
		res, err := productRepo.Delete(1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})
}
