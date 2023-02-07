package repository

// These methods can be moved inside order_repository or kept here for ease of reading

func (r Repository) CreateOrderComponents() {
	// Takes model.OrderComponent and orderId input
	// Calls DB and creates orderComponent
	// Check for error and return model.OrderComponent with ID or error
}

func (r Repository) GetOrderComponent() {
	// Takes orderID and componentID
	// Calls DB and finds particular orderComponent within a particular order ID
	// Check for error and return mode.OrderComponent or error
}

func (r Repository) GetOrderComponentsByOrderID() {
	// Takes only order ID
	// Calls DB and finds all orderComponents by order ID
	// Check for error and return []mode.OrderComponent or error
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
