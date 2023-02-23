package repository

import (
	"errors"
	"go-rest-api-template/model"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// This file only interacts with DB
// Only interacts with emailInvitee table that too

type EmailInviteeRepository interface {
	CreateEmailInvitee(model.EmailInvite) (model.EmailInvite, error)
	GetEmailInviteeById(uuid.UUID) (model.EmailInvite, error)
	UpdateEmailInviteeById()
	DeleteEmailInviteeById()
}

func NewEmailInviteeRepository(db *gorm.DB, logger log.Logger) EmailInviteeRepository {
	return Repository{db, logger}
}

func (r Repository) CreateEmailInvitee(emailInvitee model.EmailInvite) (model.EmailInvite, error) {
	// Calls DB and creates emailInvitee
	result := r.db.Create(emailInvitee)
	// Check for error
	if result.Error != nil {
		r.logger.Println("Error while saving email invitee to DB. Err: ", result.Error)
		return model.EmailInvite{}, result.Error
	}

	return emailInvitee, nil
}

func (r Repository) GetEmailInviteeById(emailInviteeId uuid.UUID) (model.EmailInvite, error) {
	// Calls DB and finds emailInvitee
	emailInvitee := model.EmailInvite{ID: emailInviteeId}
	result := r.db.Find(&emailInvitee)

	// Check for error and return mode.emailInvitee or error
	if result.Error != nil {
		// Check if error is due to record not found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			r.logger.Println("No email invitee found for ID: ", emailInviteeId)
		}

		r.logger.Println("Error while fetching email invitee from DB. Err: ", result.Error)
		return model.EmailInvite{}, result.Error
	}

	return emailInvitee, nil
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
