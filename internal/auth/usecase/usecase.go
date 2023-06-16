package usecase

import "boilerplate-clean-arch/application/domains/auth"

type authUseCase struct {
	userRepo auth.Repository
}

// Constructor
func NewAuthUseCase(userRepo auth.Repository) auth.UseCase {
	return &authUseCase{userRepo}
}
