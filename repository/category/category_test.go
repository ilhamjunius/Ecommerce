package category

import (
	"ecommerce/configs"
	"ecommerce/entities"
	"ecommerce/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoryRepo(t *testing.T) {
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

	categoryRepo := NewCategoryRepo(db)

	t.Run("Insert Category into Database", func(t *testing.T) {
		var mockCategory entities.Category
		mockCategory.CategoryType = "mainan"

		res, err := categoryRepo.Create(mockCategory)
		assert.Nil(t, err)
		assert.Equal(t, mockCategory.CategoryType, res.CategoryType)
		assert.Equal(t, 1, int(res.ID))
	})

	t.Run("Select Categories from Database", func(t *testing.T) {
		res, err := categoryRepo.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})
	t.Run("Select Category from Database", func(t *testing.T) {
		res, err := categoryRepo.Get(1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})
	t.Run("Update Category ", func(t *testing.T) {
		var mockCategory entities.Category
		mockCategory.CategoryType = "elektronik"
		res, err := categoryRepo.Update(mockCategory, 1)
		assert.Nil(t, err)
		assert.Equal(t, mockCategory.CategoryType, res.CategoryType)
	})
	t.Run("Delete Category", func(t *testing.T) {
		res, err := categoryRepo.Delete(1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})
	//TEST ERROR
	db.Migrator().DropTable(&entities.ShoppingCart{})
	db.Migrator().DropTable(&entities.Order{})
	db.Migrator().DropTable(&entities.Product{})
	db.Migrator().DropTable(&entities.Category{})
	db.Migrator().DropTable(&entities.User{})

	t.Run("Insert Category into Database", func(t *testing.T) {
		var mockCategory entities.Category
		mockCategory.CategoryType = "mainan"

		_, err := categoryRepo.Create(mockCategory)
		assert.NotNil(t, err)
	})
	t.Run("Error Select CategoryById from Database", func(t *testing.T) {
		_, err := categoryRepo.GetAll()
		assert.NotNil(t, err)
	})
	t.Run("Error Select CategoryById from Database", func(t *testing.T) {
		_, err := categoryRepo.Get(1)
		assert.NotNil(t, err)
	})
	t.Run("Error Update User ", func(t *testing.T) {
		var mockCategory entities.Category
		mockCategory.CategoryType = "elektronik"
		_, err := categoryRepo.Update(mockCategory, 2)
		assert.NotNil(t, err)
	})
	t.Run("Error Delete Category", func(t *testing.T) {
		_, err := categoryRepo.Delete(1)
		assert.NotNil(t, err)
	})

}
