package middleware

import (
	"boilerplate-clean-arch/config"
	"boilerplate-clean-arch/internal/auth"
)

// Middleware manager
type MiddlewareManager struct {
	cfg    *config.Config
	authUC auth.UseCase
}

// Middleware manager constructor
func NewMiddlewareManager(cfg *config.Config, authUC auth.UseCase) *MiddlewareManager {
	return &MiddlewareManager{
		cfg:    cfg,
		authUC: authUC,
	}
}
