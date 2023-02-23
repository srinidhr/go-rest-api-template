package service

import (
	"errors"
	"fmt"
	"go-rest-api-template/constants"
	"go-rest-api-template/model"
	"go-rest-api-template/repository"
	"go-rest-api-template/utils"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Service encapsulates usecase logic for orders.
type OrderService interface {
	SaveOrder(model.Order) (model.Order, model.AppError)
	GetOrderById(uuid.UUID) (model.Order, error)
	VerifyOrderEmailIsActive(uuid.UUID) (model.Order, error)
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

func (s orderService) SaveOrder(orderPayload model.Order) (model.Order, model.AppError) {
	// Validate orderComponents array is not null
	if len(orderPayload.OrderComps) == 0 {
		s.logger.Println("No order components found in received order payload")
		return model.Order{}, model.AppError{Error: constants.ErrOrderComponentsNotFound, Code: http.StatusBadRequest, Message: constants.ErrOrderComponentsNotFound.Error()}
	}

	// TODO later: Make sure each component has an approved ID from enums

	/**
		Call a function like how we have getSubscriptionDate()
		It should return UTC formatted date which then we can update the orderDetails.RequestedStartDate
		It should also return subscriptionDate as second response to use in email
		It should return correct error if any
	**/
	utcFormattedTime, subscriptionDate, err := parseOrderSubscriptionDate(orderPayload, s.logger)
	if err != nil {
		return model.Order{}, model.AppError{Error: err, Code: http.StatusBadRequest, Message: constants.ErrTimeParseError.Error()}
	}

	// Update the payload with utc formatted subscription start date
	orderPayload.RequestedStartDate = utcFormattedTime.String()

	// Save order details
	result, err := s.orderRepository.CreateOrder(orderPayload)
	// Log and return error if present
	if err != nil {
		return model.Order{}, model.AppError{Error: err, Code: http.StatusInternalServerError, Message: constants.ErrSaveOrderError.Error()}
	}

	// Update each order component with the order ID
	for i := range orderPayload.OrderComps {
		orderPayload.OrderComps[i].OrderID = result.ID
	}

	// Save orderComponents by passing orderComponents array and check for error
	if err := s.orderRepository.CreateOrderComponents(orderPayload.OrderComps); err != nil {
		// TODO: Invoke delete order by ID flow
		return model.Order{}, model.AppError{Error: err, Code: http.StatusInternalServerError, Message: constants.ErrOrderComponentsNotFound.Error()}
	}

	// Create new uuid for emailInvitee ID
	emailUUID := uuid.New()

	// Create a record of the latest active emailUUID
	emailInviteDetails := model.EmailInvite{
		ID:                       emailUUID,
		OrderID:                  orderPayload.ID,
		PrevEmailInvite:          emailUUID, // refer to itself for the first entry for a particular order
		ProvisioningContactEmail: orderPayload.ProvisioningContactEmail,
		IsActive:                 true,
		CreatedAt:                orderPayload.CreatedAt,
	}

	emailInviteDetails, err = s.emailInviteeRepository.CreateEmailInvitee(emailInviteDetails)
	if err != nil {
		// TODO: Invoke delete order by ID flow
		return model.Order{}, model.AppError{Error: err, Code: http.StatusInternalServerError, Message: constants.ErrSaveEmailInviteeError.Error()}
	}

	// Construct model.EmailInvitee object with all fields and save to DB
	sendProvisioningEmail := model.ProvisioningEmail{
		SubscriptionSKU:          orderPayload.SubscriptionSKU,
		SubscriptionID:           orderPayload.SubscriptionId,
		RedirectURL:              os.Getenv("PHANES_PROVISIONING_URL") + "/" + emailInviteDetails.ID.String(),
		ProvisioningContactEmail: orderPayload.ProvisioningContactEmail,
		EndCustomerName:          orderPayload.EndCustomerName,
		SenderEmail:              os.Getenv("PHANES_SENDER_EMAIL"),
		SubscriptionStartDate:    subscriptionDate,
		DelegateURL:              os.Getenv("PHANES_PROVISIONING_URL") + "/" + emailInviteDetails.ID.String(),
	}

	// Parses order payload and formats email template with required values
	messageBody, subject, err := s.emailService.GetMessageAndSubjectBody(sendProvisioningEmail, false)
	if err != nil {
		s.logger.Println("Error occurred fetching the message and subject body. Err - ", err)
		// TODO: Invoke delete order by ID flow
		return model.Order{}, model.AppError{Error: err, Code: http.StatusInternalServerError, Message: constants.ErrEmailTemplateError.Error()}
	}

	recipients := []*string{&sendProvisioningEmail.ProvisioningContactEmail}

	// Get Aws SES session
	awsRegion := os.Getenv("PHANES_AWS_REGION")
	sesSession, err := utils.GetAwsSesSession(awsRegion)
	if err != nil {
		s.logger.Println("Error occurred while creating aws session. Err - ", err)
		// TODO: Invoke delete order by ID flow
		return model.Order{}, model.AppError{Error: err, Code: http.StatusInternalServerError, Message: constants.ErrAwsSessionError.Error()}
	}

	if emailError := s.emailService.SendEmailUsingAwsSes(sesSession, messageBody, subject, recipients, sendProvisioningEmail.SenderEmail); emailError != nil {
		// TODO: Invoke delete order by ID flow
		return model.Order{}, model.AppError{Error: err, Code: http.StatusInternalServerError, Message: constants.ErrEmailSendError.Error()}
	}

	/**
		If no error, then return the following:
		- Order details with unique ID and createdAt
		- Order details should have all the orderComponents
	**/
	return orderPayload, model.AppError{}
}

func (s orderService) GetOrderById(orderId uuid.UUID) (model.Order, error) {
	// Call repository layer to fetch all orders
	result, err := s.orderRepository.GetOrderById(orderId)
	// Check for errors
	if err != nil {
		return model.Order{}, err
	}

	// Call repository to fetch all order components for particular orderId
	orderComponents, err := s.orderRepository.GetOrderComponentsByOrderID(orderId)
	// Check for errors
	if err != nil {
		return model.Order{}, err
	}

	// Create complete model.Order for response and return model.Order
	result.OrderComps = orderComponents
	return result, nil
}

func (s orderService) VerifyOrderEmailIsActive(emailUUID uuid.UUID) (model.Order, error) {
	// Fetch emailInvitee by ID
	result, err := s.emailInviteeRepository.GetEmailInviteeById(emailUUID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.Order{}, errors.New("bad_request")
		}

		return model.Order{}, err
	}

	// Check if UUID is active
	if !result.IsActive {
		s.logger.Println("emailUUID received is not active")
		return model.Order{}, errors.New("bad_request")
	}

	// Fetch order using orderId and this service layer function itself
	return s.GetOrderById(result.OrderID)
}

func parseOrderSubscriptionDate(orderDetails model.Order, logger log.Logger) (time.Time, string, error) {
	// Parse requested start date in payload
	parsedTime, timeParseErr := time.Parse(time.RFC3339, orderDetails.RequestedStartDate) //https://pkg.go.dev/time
	if timeParseErr != nil {
		logger.Println("Error while parsing requestedStartDate. Err - ", timeParseErr)
		return time.Time{}, "", timeParseErr
	}

	if parsedTime.Location() != time.UTC {
		logger.Println("Received time is not in UTC format, converting...")
		//For RFC3339 it loses the seconds in Decimal Places
		parsedTime.UTC().Format(time.RFC3339Nano) //Used RFC3339Nano to keep fraction seconds,
	}

	// Format subscriptionDate to send in email
	subscriptionDate := "XX/XX/XXXX"
	parsedDate, dateParseErr := time.Parse(time.RFC3339, orderDetails.RequestedStartDate)
	year, month, day := parsedDate.Date()

	if dateParseErr == nil {
		subscriptionDate = fmt.Sprintf("%02d/%02d/%d", int(month), day, year)
	}
	return parsedTime, subscriptionDate, nil
}
