package middleware

import (
	"go-clean-arch/config"
	authRepo "go-clean-arch/internal/auth/repository"
)

// Middleware manager
type MiddlewareManager struct {
	cfg      *config.Config
	authRepo authRepo.IRepository
}

// Middleware manager constructor
func NewMiddlewareManager(cfg *config.Config, authRepo authRepo.IRepository) *MiddlewareManager {
	return &MiddlewareManager{
		cfg:      cfg,
		authRepo: authRepo,
	}
}
