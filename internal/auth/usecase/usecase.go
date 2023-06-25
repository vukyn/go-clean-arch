package usecase

import (
	"boilerplate-clean-arch/config"
	"boilerplate-clean-arch/internal/auth"
	"boilerplate-clean-arch/internal/session"
)

type authUseCase struct {
	cfg      *config.Config
	userRepo auth.Repository
	sessRepo session.SessRepository
}

// Constructor
func NewAuthUseCase(cfg *config.Config, userRepo auth.Repository, sessRepo session.SessRepository) auth.UseCase {
	return &authUseCase{
		cfg: 	cfg,
		userRepo: userRepo,
		sessRepo: sessRepo,
	}
}
