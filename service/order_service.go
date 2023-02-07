package service

import (
	"go-rest-api-template/repository"
	"log"
)

// Service encapsulates usecase logic for orders.
type OrderService interface {
	SaveOrder()
	GetOrderById()
	VerifyOrderEmailIsActive()
}

type orderService struct {
	orderRepository        repository.OrderRepository
	emailInviteeRepository repository.EmailInviteeRepository
	emailService           EmailService
	logger                 log.Logger
}

// NewService creates a new order service.
func NewOrderService(orderRepository repository.OrderRepository,
	emailInviteeRepository repository.EmailInviteeRepository,
	emailService EmailService,
	logger log.Logger) OrderService {
	return orderService{orderRepository, emailInviteeRepository, emailService, logger}
}

func (s orderService) SaveOrder() {
	// Take model.Order as input
	// Validate orderComponents array is not null

	// TODO later: Make sure each component has an approved ID from enums

	/**
		Call a function like how we have getSubscriptionDate()
		Name the function as "parseOrderSubscriptionDate()"
		It should return UTC formatted date which then we can update the orderDetails.RequestedStartDate
		It should also return subscriptionDate as second response to use in email
		It should return correct error if any

		Return error if present
		Update orderDetails.RequestedStartDate with formatted date
		Keep subscriptionDate in "XX/XX/XXXX" for later
	**/

	// Copy over orderComponents to new array
	// Remove orderComponents from original payload

	// Create unique Order ID
	// Add createdAt timestamp

	// Save order details
	s.orderRepository.CreateOrder()
	// Log and return error if present

	// Save orderComponents by passing orderComponents array and orderId
	s.orderRepository.CreateOrderComponents()

	// Create new uuid for emailInvitee ID
	// Construct model.EmailInvitee object with all fields and save to DB
	s.emailInviteeRepository.CreateEmailInvitee()

	// Construct model.ProvisioningEmail
	// Get messageBody and subject
	s.emailService.GetMessageBodyAndSubject()

	// Fetch list of receipients and also create AWS Session

	// Call emailService to send email by injecting SES session
	s.emailService.SendEmailUsingAwsSes()

	/**
		If no error, then return the following:
		- Order details with unique ID and createdAt
		- Order details should have all the orderComponents
		- Each order components do not need to have orderId in response as it is used only for internal relational DB operations
	**/

	/**
		If error,
		return error while sending email

		TODO later: We need to remove from DB whenver any subsequent step fails
		We should not be storing any orders which have not gone through the entire pipeline and email has not been sent successfully
	**/
}

func (s orderService) GetOrderById() {
	// Take orderId as input

	// Call repository layer to fetch all orders
	s.orderRepository.GetOrderById()
	// Handle any errors

	// Call repository to fetch all order components for particular orderId
	s.orderRepository.GetOrderComponentsByOrderID()
	// Handle any errors

	// Create complete model.Order for response and return model.Order
}

func (s orderService) VerifyOrderEmailIsActive() {
	// Takes UUID as input

	// Fetch emailInvitee by ID
	s.emailInviteeRepository.GetEmailInviteeById()

	// Check if UUID is active

	// Fetch order using orderId and this service layer function itself
	s.GetOrderById()

	// Return order details itself
	// Maybe order quantity in each component can be omitted if not required by UX
}
