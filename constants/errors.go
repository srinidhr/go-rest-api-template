package constants

import "errors"

var (
	ErrOrderComponentsNotFound  = errors.New("no order components found")
	ErrorUUIDNotActive          = errors.New("UUID not active")
	ErrTimeParseError           = errors.New("error in parsing requestedStartDate to UTC")
	ErrSaveOrderError           = errors.New("error while saving order to DB")
	ErrSaveOrderComponentsError = errors.New("error while saving order components to DB")
	ErrSaveEmailInviteeError    = errors.New("error while saving email invitee to DB")
	ErrEmailTemplateError       = errors.New("error while constructing email template")
	ErrAwsSessionError          = errors.New("error in creating AWS session")
	ErrEmailSendError           = errors.New("error while trying to send email for provisioning")
)
