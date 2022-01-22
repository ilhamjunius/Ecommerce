package order

import (
	"ecommerce/entities"

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
		return nil, err
	}
	return orders, nil
}

func (or *OrderRepository) Get(orderId, userId int) (entities.Order, error) {
	order := entities.Order{}
	if err := or.db.Find(&order, "id=? AND user_id=?", orderId, userId).Error; err != nil {
		return order, err
	}
	return order, nil
}

func (or *OrderRepository) Create(newOrder entities.Order) (entities.Order, error) {
	shoppingcart := []entities.ShoppingCart{}
	if err := or.db.Find(&shoppingcart, "order_id=?", newOrder).Error; err != nil {
		return newOrder, err
	}

	total := 0
	for a := 0; a < len(shoppingcart); a++ {
		total += shoppingcart[a].Subtotal
	}

	newOrder.Total = total

	if err := or.db.Save(&newOrder).Error; err != nil {
		return newOrder, err
	}
	return newOrder, nil
}

func (or *OrderRepository) Cancel(orderId, userId int) (entities.Order, error) {
	order := entities.Order{}
	if err := or.db.First(&order, "id=? AND user_id=?", orderId, userId).Error; err != nil {
		return order, err
	}

	order.InvoiceID = "Cancel"
	order.PaymentLink = "Cancel"
	order.Status = "Cancel"

	or.db.Save(&order)

	return order, nil
}

func (or *OrderRepository) Pay(invoiceId, paymentLink string, orderId, userId int) (entities.Order, error) {
	order := entities.Order{}
	if err := or.db.Find(&order, "id=? AND user_id=?", orderId, userId).Error; err != nil {
		return order, err
	}

	order.InvoiceID = invoiceId
	order.PaymentLink = paymentLink
	order.Status = "Waiting for Payment"

	or.db.Save(&order)

	return order, nil
}

func (or *OrderRepository) Check(orderId, userId int) (entities.Order, error) {
	order := entities.Order{}
	if err := or.db.Find(&order, "id=? AND user_id=?", orderId, userId).Error; err != nil {
		return order, err
	}

	order.Status = "Paid"

	or.db.Save(&order)

	return order, nil
}
