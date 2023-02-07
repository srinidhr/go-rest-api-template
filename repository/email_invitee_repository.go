package repository

import (
	"log"

	"gorm.io/gorm"
)

// This file only interacts with DB
// Only interacts with emailInvitee table that too

type EmailInviteeRepository interface {
	CreateEmailInvitee()
	GetEmailInviteeById()
	UpdateEmailInviteeById()
	DeleteEmailInviteeById()
}

func NewEmailInviteeRepository(db *gorm.DB, logger log.Logger) EmailInviteeRepository {
	return Repository{db, logger}
}

func (r Repository) CreateEmailInvitee() {
	// Takes model.EmailInvitee has input
	// Create unique EmailInvitee ID
	// Calls DB and creates emailInvitee
	// Check for error and return model.EmailInvitee with ID or error
}

func (r Repository) GetEmailInviteeById() {
	// Takes only emailInvitee ID
	// Calls DB and finds emailInvitee
	// Check for error and return mode.emailInvitee or error
}

func (r Repository) UpdateEmailInviteeById() {
	// Takes emailInvitee ID and new model.EmailInvitee with updated fields as input
	// Calls DB to update fields by ID
	// Check for error and return model.EmailInvitee or error
}

func (r Repository) DeleteEmailInviteeById() {
	// Takes emailInvitee ID as input
	// Calls DB to delete emailInvitee by ID
	// Return is only if error happens
}
