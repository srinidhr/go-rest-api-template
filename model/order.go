package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID                       uuid.UUID        `gorm:"type:uuid;primary_key;"`
	ProvisioningContactEmail string           `json:"provisioningContactEmail"`
	EndCustomerName          string           `json:"endCustomerName"`
	SubscriptionId           string           `json:"subscriptionId"`
	ProvisioningAction       string           `json:"action"`
	SubscriptionSKU          string           `json:"subscriptionSKU"`
	OrderComps               []OrderComponent `json:"orderComponents" gorm:"-"` // "-" tag ignores this field when writing to DB
	CreatedAt                time.Time        `json:"createdAt"`
	RequestedStartDate       string           `json:"requestedStartDate"`
}

type OrderComponent struct {
	OrderID         uuid.UUID `json:"orderID"`
	ComponentID     string    `json:"componentID"`
	QuantityOrdered uint32    `json:"quantityOrdered"`
}

type EmailInvite struct {
	ID                       uuid.UUID `gorm:"type:uuid;primary_key;"`
	OrderID                  uuid.UUID
	PrevEmailInvite          uuid.UUID
	ProvisioningContactEmail string
	IsActive                 bool
	CreatedAt                time.Time
}
