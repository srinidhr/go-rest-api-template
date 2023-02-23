package repository

import (
	"errors"
	"go-rest-api-template/model"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// This file only interacts with DB
// Only interacts with orders related tables that too

type OrderRepository interface {
	// orders table functions
	CreateOrder(model.Order) (model.Order, error)
	GetOrderById(uuid.UUID) (model.Order, error)
	GetAllOrders() ([]model.Order, error)
	UpdateOrderById()
	DeleteOrderById(uuid.UUID) error

	// order_components table functions
	CreateOrderComponents([]model.OrderComponent) error
	GetOrderComponent()
	GetOrderComponentsByOrderID(uuid.UUID) ([]model.OrderComponent, error)
	UpdateOrderComponentsByOrderId()
	DeleteOrderComponentsByOrderId()
}

func NewOrderRepository(db *gorm.DB, logger log.Logger) OrderRepository {
	return Repository{db, logger}
}

func (r Repository) CreateOrder(order model.Order) (model.Order, error) {
	// Create unique Order ID
	order.ID = uuid.New()
	// Add createdAt timestamp
	order.CreatedAt = time.Now().UTC()

	// Calls DB and creates order
	result := r.db.Omit("OrderComps").Create(order)

	// Check for error
	if result.Error != nil {
		r.logger.Println("Order creation failed. Error: ", result.Error)
		return model.Order{}, result.Error
	}

	r.logger.Println("Order created successfully. ID: ", order.ID)
	return order, nil
}

func (r Repository) GetOrderById(orderId uuid.UUID) (model.Order, error) {
	// Calls DB and finds order
	orderDetails := model.Order{ID: orderId}
	result := r.db.Find(&orderDetails)

	if result.Error != nil {
		// Check if error is due to record not found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			r.logger.Println("No order found for ID: ", orderId)
		}

		// Return error
		r.logger.Println("Error while getting order by ID: ", orderId)
		return model.Order{}, result.Error
	}

	// Return orderDetails if found
	r.logger.Println("Order found for given ID: ", orderId)
	return orderDetails, nil
}

func (r Repository) GetAllOrders() ([]model.Order, error) {
	// Calls DB and finds all order
	allOrders := []model.Order{}
	result := r.db.Find(&allOrders)

	// Check for error
	if result.Error != nil {
		r.logger.Println("Error getting all orders. Err: ", result.Error)
		return []model.Order{}, result.Error
	}

	// Log success and return all orders
	r.logger.Println("All orders found successfully. Length: ", result.RowsAffected)
	return allOrders, nil
}

func (r Repository) UpdateOrderById() {
	// Takes order ID and new model.Order with updated fields as input
	// Calls DB to update fields by ID
	// Check for error and return model.Order or error
}

func (r Repository) DeleteOrderById(orderId uuid.UUID) error {
	// Calls DB to delete order by ID
	result := r.db.Delete(&model.Order{ID: orderId})

	// Return only if error
	return result.Error
}
