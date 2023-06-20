package usecase

import (
	"boilerplate-clean-arch/config"
	"boilerplate-clean-arch/internal/auth"
)

type authUseCase struct {
	cfg      *config.Config
	userRepo auth.Repository
}

// Constructor
func NewAuthUseCase(cfg *config.Config, userRepo auth.Repository) auth.UseCase {
	return &authUseCase{cfg, userRepo}
}
