package product

import (
	"ecommerce/entities"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (pr *ProductRepository) GetAll(keyword string) ([]entities.Product, error) {
	Products := []entities.Product{}
	// if err := pr.db.Find(&Products).Error; err != nil {
	// 	return nil, err
	// }
	// keyword := "sapu%"
	if err := pr.db.Where("Name like ?", keyword).Find(&Products).Error; err != nil {
		return Products, err
	}
	return Products, nil
}
func (pr *ProductRepository) FilterProduct(keyword, category string) ([]entities.Product, error) {
	Products := []entities.Product{}
	keywordd := "sapu"
	if err := pr.db.Where("Name like ?", keywordd+"%").Find(&Products).Error; err != nil {
		return Products, err
	}
	return Products, nil
}
func (pr *ProductRepository) Get(productId int) (entities.Product, error) {
	Product := entities.Product{}
	if err := pr.db.Find(&Product, productId).Error; err != nil {
		return Product, err
	}
	return Product, nil
}

func (pr *ProductRepository) Create(newProduct entities.Product) (entities.Product, error) {

	if err := pr.db.Save(&newProduct).Error; err != nil {
		return newProduct, err
	}
	return newProduct, nil
}

func (pr *ProductRepository) Update(newProduct entities.Product, productId int) (entities.Product, error) {
	product := entities.Product{}
	if err := pr.db.First(&product, "id=?", productId).Error; err != nil {
		return newProduct, err
	}

	product.Name = newProduct.Name
	product.Price = newProduct.Price
	product.Stock = newProduct.Stock
	product.CategoryID = newProduct.CategoryID
	product.Description = newProduct.Description

	pr.db.Save(&product)

	return newProduct, nil
}

func (pr *ProductRepository) Delete(productId int) (entities.Product, error) {
	product := entities.Product{}

	if err := pr.db.Delete(&product, productId).Error; err != nil {
		return product, err
	}

	return product, nil
}
