package order

import (
	"ecommerce/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (or *OrderRepository) GetAll(userId int) ([]entities.Order, error) {
	orders := []entities.Order{}
	if err := or.db.Where("user_id=?", userId).Find(&orders).Error; err != nil {
		log.Warn("Found database error", err)
		return nil, err
	}
	return orders, nil
}

func (or *OrderRepository) Get(orderId, userId int) (entities.Order, error) {
	order := entities.Order{}
	if err := or.db.Find(&order, "order_id=? AND user_id=?", orderId, userId).Error; err != nil {
		log.Warn("Found database error", err)
		return order, err
	}
	return order, nil
}

func (or *OrderRepository) Create(newOrder entities.Order) (entities.Order, error) {
	if err := or.db.Save(&newOrder).Error; err != nil {
		return newOrder, err
	}
	return newOrder, nil
}

func (or *OrderRepository) Update(newOrder entities.Order, orderId, userId int) (entities.Order, error) {
	order := entities.Order{}
	if err := or.db.Find(&order, "order_id=? AND user_id=?", orderId, userId).Error; err != nil {
		return order, err
	}

	order = newOrder

	if err := or.db.Save(&order).Error; err != nil {
		return newOrder, err
	}

	return newOrder, nil
}

func (or *OrderRepository) Delete(orderId, userId int) (entities.Order, error) {
	order := entities.Order{}
	if err := or.db.Find(&order, "order_id=? AND user_id=?", orderId, userId).Error; err != nil {
		return order, err
	}
	if err := or.db.Delete(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}
