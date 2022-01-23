package category

import (
	"ecommerce/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (cr *CategoryRepository) GetAll() ([]entities.Category, error) {
	categories := []entities.Category{}
	if err := cr.db.Find(&categories).Error; err != nil {
		log.Warn("Found database error", err)
		return nil, err
	}

	return categories, nil
}

func (cr *CategoryRepository) Get(categoryId int) (entities.Category, error) {
	category := entities.Category{}
	if err := cr.db.First(&category, categoryId).Error; err != nil {
		log.Warn("Found database error", err)
		return category, err
	}
	return category, nil
}

func (cr *CategoryRepository) Create(newCategory entities.Category) (entities.Category, error) {
	if err := cr.db.Save(&newCategory).Error; err != nil {
		return newCategory, err
	}
	return newCategory, nil
}

func (cr *CategoryRepository) Update(newCategory entities.Category, categoryId int) (entities.Category, error) {
	category := entities.Category{}
	if err := cr.db.First(&category, "id=?", categoryId).Error; err != nil {
		return newCategory, err
	}
	cr.db.Model(&category).Updates(newCategory)

	return newCategory, nil
}

func (cr *CategoryRepository) Delete(categoryId int) (entities.Category, error) {
	category := entities.Category{}
	if err := cr.db.Find(&category, "id=?", categoryId).Error; err != nil {
		return category, err
	}
	cr.db.Delete(&category)
	return category, nil
}
