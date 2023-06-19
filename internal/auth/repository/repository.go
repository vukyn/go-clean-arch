package repository

import (
	"boilerplate-clean-arch/internal/auth"

	"gorm.io/gorm"
)

// Auth repository
type authRepo struct {
	db *gorm.DB
}

// Constructor
func NewAuthRepository(db *gorm.DB) auth.Repository {
	return &authRepo{
		db: db,
	}
}