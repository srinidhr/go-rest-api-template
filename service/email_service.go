package service

import (
	"bufio"
	"go-rest-api-template/model"
	"io"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
)

// Create a function to equate to file open function
var fileOpen = os.Open

type EmailService interface {
	SendEmailUsingAwsSes(sesiface.SESAPI, string, string, []*string, string) error
	GetMessageAndSubjectBody(model.ProvisioningEmail, bool) (string, string, error)
}

type emailService struct {
	logger log.Logger
}

func NewEmailService(logger log.Logger) EmailService {
	return emailService{logger}
}

// SendEmailUsingAwsSes method is used to email the end user
// using AWS SES service.
func (s emailService) SendEmailUsingAwsSes(sesSession sesiface.SESAPI, messageBody string, subject string,
	recipients []*string, senderEmail string) error {

	input := &ses.SendEmailInput{

		Destination: &ses.Destination{
			ToAddresses: recipients,
		},

		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(messageBody),
				},
			},

			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(subject),
			},
		},

		Source: aws.String(senderEmail),
	}

	outputEmail, err := sesSession.SendEmail(input)
	s.logger.Println(outputEmail.MessageId)
	if err != nil {
		s.logger.Println("Error sending mail - ", err)
		return err
	}

	return nil
}

// getMessageAndSubjectBody method is used to create the message body and subject line
// for the email. It refers to a template file and reads the html content from it.
// Further, replaces the keywords with dynamic data like subscriptionID, subscriptionSKU etc
func (s emailService) GetMessageAndSubjectBody(awsSesProvisioningEmail model.ProvisioningEmail, isDelegateConfirmationEmail bool) (string, string, error) {

	var fileName string
	if isDelegateConfirmationEmail {
		fileName = "delegate_email_template.html"
	} else {
		fileName = "email_template.html"
	}

	f, err := fileOpen(fileName)
	if err != nil {
		s.logger.Println("Error opening email template path. Err - ", err)
		return "", "", err
	}

	reader := bufio.NewReader(f)
	buf := make([]byte, 1024)
	messageBody := ""
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				s.logger.Println(err)
			}
			break
		}
		messageBody += string(buf[0:n])
	}
	messageBody = ReplaceSesMessageBodyWithActualValues(awsSesProvisioningEmail, messageBody, isDelegateConfirmationEmail)
	subjectBody := "Cisco Security Cloud <noreply@cisco.com>"
	return messageBody, subjectBody, nil
}

func ReplaceSesMessageBodyWithActualValues(awsSesProvisioningEmail model.ProvisioningEmail, messageBody string, isDelegateConfirmationEmail bool) string {
	if isDelegateConfirmationEmail {
		messageBody = strings.ReplaceAll(messageBody, "{{DELEGATE_TO_EMAIL}}",
			awsSesProvisioningEmail.DelegatedEmail)

		messageBody = strings.ReplaceAll(messageBody, "{{SUPPORT_EMAIL}}", awsSesProvisioningEmail.SupportEmail)
	} else {
		messageBody = strings.ReplaceAll(messageBody, "{{SUBSCRIPTION_SKU}}",
			awsSesProvisioningEmail.SubscriptionSKU)

		messageBody = strings.Replace(messageBody, "{{SUBSCRIPTION_ID}}",
			awsSesProvisioningEmail.SubscriptionID, 1)

		messageBody = strings.Replace(messageBody, "{{SUBSCRIPTION_START_DATE}}",
			awsSesProvisioningEmail.SubscriptionStartDate, 1)

		messageBody = strings.Replace(messageBody, "{{REDIRECT_URL}}",
			awsSesProvisioningEmail.RedirectURL, 1)

		messageBody = strings.Replace(messageBody, "{{DELEGATE_URL}}",
			awsSesProvisioningEmail.DelegateURL, 1)
	}
	return messageBody
}
