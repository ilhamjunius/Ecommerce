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

// func (or *OrderRepository) Create(newOrder entities.Order, orderId int) (entities.Order, error) {
// 	order := entities.Order{}
// 	or.db.Table("shoppingcart").Select("order_id,user_id,sum(subtotal) as total").Where("order_id LIKE ?", orderId).Group("order_id").First(&order)
// 	//XENDIT LOGIN
// 	xenditApiKey := "eG5kX2RldmVsb3BtZW50X3REdlpkNkJycEh3RVN1dU1ZWGpWaFZYdk1yZW9BcHNWbXd4VFl2cDJIUThIbVRvWFFiNTQydFA0QUlLamYxbzgK:"
// 	data := invoice.CreateParams{
// 		ExternalID: order.ID,
// 		Amount:     order.Total,
// 	}
// 	response, err := invoice.Create(&data)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	order.PaymentId = response
// 	return newOrder, nil
// }

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
