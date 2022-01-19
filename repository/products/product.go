package product

import (
	"ecommerce/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (pr *ProductRepository) GetAll() ([]entities.Product, error) {
	Products := []entities.Product{}
	if err := pr.db.Find(&Products).Error; err != nil {
		log.Warn("Found database error", err)
		return nil, err
	}
	return Products, nil
}

func (pr *ProductRepository) Create(newProduct entities.Product) (entities.Product, error) {

	if err := pr.db.Save(&newProduct).Error; err != nil {
		log.Warn("Found database error", err)
		return newProduct, err
	}
	return newProduct, nil
}

func (pr *ProductRepository) Update(newProduct entities.Product, productId int) (entities.Product, error) {
	product := entities.Product{}
	if err := pr.db.First(&product, "id=?", productId).Error; err != nil {
		return newProduct, err
	}
	err := pr.db.Model(&product).Updates(entities.Product{
		Name:        newProduct.Name,
		Price:       newProduct.Price,
		Stock:       newProduct.Stock,
		CategoryID:  newProduct.CategoryID,
		Description: newProduct.Description,
	}).Error

	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func (pr *ProductRepository) Delete(productId int) (entities.Product, error) {
	product := entities.Product{}
	if err := pr.db.Find(&product, "id=?", productId).Error; err != nil {
		return product, err
	}

	if err := pr.db.Delete(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}
