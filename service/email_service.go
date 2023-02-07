package service

import (
	"log"
)

type EmailService interface {
	SendEmailUsingAwsSes()
	GetMessageBodyAndSubject()
}

type emailService struct {
	logger log.Logger
}

func NewEmailService(logger log.Logger) EmailService {
	return emailService{logger}
}

func (s emailService) SendEmailUsingAwsSes() {
	// Take sesiface.SESAPI, messageBody, subject and list of recipients as input
	// Construct input object with source, destination and mesasge configurations
	// Call SESAPI.SendEmail() with the input
	// Return only error if exists; we do not need or do anything with message ID generated from AWS
}

func (s emailService) GetMessageBodyAndSubject() {
	// Current implementation is good; change the function name as shown here
	// Move Subject Body string to constants file
}

func ReplaceEmailMessageBodyWithActualValues() {
	// Current implementation is good
	// This can also be moved to utils.AwsUtil if that makes sense
}
