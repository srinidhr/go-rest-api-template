package repository

import (
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	db     *gorm.DB
	logger log.Logger
}
