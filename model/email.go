package model

type ProvisioningEmail struct {
	SubscriptionSKU          string
	SubscriptionID           string
	RedirectURL              string
	ProvisioningContactEmail string
	EndCustomerName          string
	SenderEmail              string
	SubscriptionStartDate    string
	DelegateURL              string
	DelegatedEmail           string
	SupportEmail             string
}
