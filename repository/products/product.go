package product

import (
	"ecommerce/entities"
	"log"
	"math"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
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

	pr.db.Model(&product).Updates(newProduct)

	return newProduct, nil
}

func (pr *ProductRepository) Delete(productId int) (entities.Product, error) {
	product := entities.Product{}

	if err := pr.db.Delete(&product, productId).Error; err != nil {
		return product, err
	}

	return product, nil
}

var total_rowsGlobal int64

func (pr *ProductRepository) Pagination(name, category string, pagination entities.Pagination) (ResultPagination, int) {

	var product []entities.Product
	totalPages, fromRow, toRow := 0, 0, 0
	var totalRows int64

	categoryJoins := entities.Category{}
	if err := pr.db.Joins("JOIN products on categories.id=products.category_id").
		Where("category_type=(?)", category).
		Find(&categoryJoins).Error; err != nil {
		log.Fatal(err)
	}
	offset := pagination.Page * pagination.Limit

	if name == "" && category == "" {

		errFind := pr.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Find(&product).Count(&totalRows).Error
		if errFind != nil {
			return ResultPagination{Error: errFind}, totalPages
		}
	} else if name == "" {

		errFind := pr.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Where("category_id=?", categoryJoins.ID).Find(&product).Count(&totalRows).Error
		if errFind != nil {
			return ResultPagination{Error: errFind}, totalPages
		}
	} else if category == "" {

		errFind := pr.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Where("name LIKE ?", "%"+name+"%").Find(&product).Count(&totalRows).Error
		if errFind != nil {
			return ResultPagination{Error: errFind}, totalPages
		}
	} else {
		errFind := pr.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Where("name LIKE ? and category_id=?", "%"+name+"%", categoryJoins.ID).Find(&product).Count(&totalRows).Error
		if errFind != nil {
			return ResultPagination{Error: errFind}, totalPages
		}
	}
	// errCount := pr.db.Model(&product).Count(&totalRows).Error
	// if errCount != nil {
	// 	return ResultPagination{Error: errCount}, totalPages
	// }
	if offset == 0 {
		total_rowsGlobal = totalRows
	}

	pagination.TotalRows = int(total_rowsGlobal)
	pagination.Rows = product

	totalRows = int64(pagination.TotalRows)
	totalPages = int(math.Ceil(float64(total_rowsGlobal)/float64(pagination.Limit))) - 1
	if pagination.Page == 0 {
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = (pagination.Page + 1) * pagination.Limit
		}
	}
	if int64(toRow) > total_rowsGlobal {
		toRow = int(total_rowsGlobal)
	}
	pagination.FromRow = fromRow
	pagination.ToRow = toRow
	return ResultPagination{Result: pagination}, totalPages

}
