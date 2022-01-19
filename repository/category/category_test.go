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

	db.Migrator().DropTable(&entities.Category{})
	db.AutoMigrate(&entities.Category{})

	categoryRepo := NewCategoryRepo(db)

	t.Run("Insert Category into Database", func(t *testing.T) {
		var mockCategory entities.Category
		mockCategory.CategoryType = "mainan"

		res, err := categoryRepo.Create(mockCategory)
		assert.Nil(t, err)
		assert.Equal(t, mockCategory.CategoryType, res.CategoryType)
		assert.Equal(t, 1, int(res.ID))
	})
	t.Run("Error Insert Category into Database", func(t *testing.T) {
		var mockCategory entities.Category
		mockCategory.CategoryType = "mainan"

		_, err := categoryRepo.Create(mockCategory)
		assert.Error(t, err)

	})

	t.Run("Select Categories from Database", func(t *testing.T) {
		res, err := categoryRepo.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})
	t.Run("Error Select Categories from Database", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Category{})
		_, err := categoryRepo.GetAll()
		assert.Error(t, err)

	})
	t.Run("Insert Category into Database", func(t *testing.T) {
		var mockCategory entities.Category
		mockCategory.CategoryType = "mainan"
		db.AutoMigrate(&entities.Category{})
		res, err := categoryRepo.Create(mockCategory)
		assert.Nil(t, err)
		assert.Equal(t, mockCategory.CategoryType, res.CategoryType)
		assert.Equal(t, 1, int(res.ID))
	})
	t.Run("Select Category from Database", func(t *testing.T) {
		res, err := categoryRepo.Get(1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})
	t.Run("Error Select CategoryById from Database", func(t *testing.T) {
		_, err := categoryRepo.Get(100)
		assert.Error(t, err)

	})

	t.Run("Update User ", func(t *testing.T) {
		var mockCategory entities.Category
		mockCategory.CategoryType = "elektronik"
		res, err := categoryRepo.Update(mockCategory, 1)
		assert.Nil(t, err)
		assert.Equal(t, mockCategory.CategoryType, res.CategoryType)
	})
	t.Run("Error Update User ", func(t *testing.T) {
		var mockCategory entities.Category
		mockCategory.CategoryType = "elektronik"
		db.Migrator().DropTable(&entities.Category{})
		_, err := categoryRepo.Update(mockCategory, 2)
		assert.Error(t, err)

	})
	db.AutoMigrate(&entities.Category{})
	t.Run("Insert Category into Database", func(t *testing.T) {
		var mockCategory entities.Category
		mockCategory.CategoryType = "mainan"

		res, err := categoryRepo.Create(mockCategory)
		assert.Nil(t, err)
		assert.Equal(t, mockCategory.CategoryType, res.CategoryType)
		assert.Equal(t, 1, int(res.ID))
	})

	t.Run("Delete User", func(t *testing.T) {
		res, err := categoryRepo.Delete(1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})
	t.Run("Error Delete User", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Category{})
		_, err := categoryRepo.Delete(1)
		assert.Error(t, err)

	})
}
