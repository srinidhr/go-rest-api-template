package repository

import (
	"log"

	"gorm.io/gorm"
)

// This file only interacts with DB
// Only interacts with orders related tables that too

type OrderRepository interface {
	// orders table functions
	CreateOrder()
	GetOrderById()
	GetAllOrders()
	UpdateOrderById()
	DeleteOrderById()

	// order_components table functions
	CreateOrderComponents()
	GetOrderComponent()
	GetOrderComponentsByOrderID()
	UpdateOrderComponentsByOrderId()
	DeleteOrderComponentsByOrderId()
}

func NewOrderRepository(db *gorm.DB, logger log.Logger) OrderRepository {
	return Repository{db, logger}
}

func (r Repository) CreateOrder() {
	// Takes model.Order has input
	// Calls DB and creates order
	// Check for error and return model.Order with ID or error
}

func (r Repository) GetOrderById() {
	// Takes only order ID
	// Calls DB and finds order
	// Check for error and return mode.Order or error
}

func (r Repository) GetAllOrders() {
	// Does not take any input
	// Calls DB and finds all order
	// Check for error and return []mode.Order or error
}

func (r Repository) UpdateOrderById() {
	// Takes order ID and new model.Order with updated fields as input
	// Calls DB to update fields by ID
	// Check for error and return model.Order or error
}

func (r Repository) DeleteOrderById() {
	// Takes order ID as input
	// Calls DB to delete order by ID
	// Return is only if error happens
}
