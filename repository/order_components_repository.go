package repository

import (
	"go-rest-api-template/model"

	"github.com/google/uuid"
)

// These methods can be moved inside order_repository or kept here for ease of reading

func (r Repository) CreateOrderComponents(orderComponents []model.OrderComponent) error {
	// Calls DB and creates orderComponent
	result := r.db.Create(orderComponents)

	// Check for error
	if result.Error != nil {
		r.logger.Println("Error while creating order components in DB. Err: ", result.Error)
		return result.Error
	}

	return nil
}

func (r Repository) GetOrderComponent() {
	// Takes orderID and componentID
	// Calls DB and finds particular orderComponent within a particular order ID
	// Check for error and return mode.OrderComponent or error
}

func (r Repository) GetOrderComponentsByOrderID(orderId uuid.UUID) ([]model.OrderComponent, error) {
	// Calls DB and finds all orderComponents by order ID
	allOrderComponents := []model.OrderComponent{}
	result := r.db.Where("orderId = ?", orderId).Find(allOrderComponents)

	// Check for error and return []mode.OrderComponent or error
	if result.Error != nil {
		r.logger.Println("Error while fetching all order components for order ID: ", orderId)
		return []model.OrderComponent{}, result.Error
	}

	return allOrderComponents, nil
}

func (r Repository) UpdateOrderComponentsByOrderId() {
	// Takes orderComponent ID and new model.OrderComponent with updated fields as input
	// Calls DB to update fields by ID
	// Check for error and return model.OrderComponent or error
}

func (r Repository) DeleteOrderComponentsByOrderId() {
	// Takes order ID as input
	// Calls DB to delete all orderComponents associated to order ID
	// Return is only if error happens
}
