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
		db.Migrator().DropTable(&entities.Product{})
		_, err := productRepo.Create(mockProduct)
		assert.Error(t, err)

	})
	db.AutoMigrate(&entities.Product{})
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
	db.Migrator().DropTable(&entities.Product{})
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
	db.AutoMigrate(&entities.Product{})
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

	// t.Run("Insert User into Database", func(t *testing.T) {
	// 	password := sha256.Sum256([]byte("andrew123"))
	// 	var mockInsertUser entities.User
	// 	mockInsertUser.Email = "andrew@yahoo.com"
	// 	mockInsertUser.Password = password[:]
	// 	mockInsertUser.Name = "andrew"
	// 	mockInsertUser.HandphoneNumber = "0123456789"
	// 	mockInsertUser.Role = "admin"

	// 	res, err := userRepo.Create(mockInsertUser)
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, mockInsertUser.Name, res.Name)
	// 	assert.Equal(t, 2, int(res.ID))
	// })

	// t.Run("Update User ", func(t *testing.T) {
	// 	password := sha256.Sum256([]byte("ilham123"))
	// 	var mockUpdateUser entities.User
	// 	mockUpdateUser.Email = "ilham@yahoo.com"
	// 	mockUpdateUser.Password = password[:]
	// 	mockUpdateUser.Name = "ilham"
	// 	mockUpdateUser.HandphoneNumber = "987654321"
	// 	mockUpdateUser.Role = "pembeli"

	// 	res, err := userRepo.Create(mockUpdateUser)
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, mockUpdateUser.Name, res.Name)
	// })
}
