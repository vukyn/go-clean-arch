package repository

import (
	"boilerplate-clean-arch/application/domains/auth"

	"gorm.io/gorm"
)

// Auth repository
type authRepo struct {
	db *gorm.DB
}

// Constructor
func NewAuthRepo(db *gorm.DB) auth.Repository {
	return &authRepo{
		db: db,
	}
}