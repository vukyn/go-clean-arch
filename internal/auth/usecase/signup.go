package usecase

import (
	"boilerplate-clean-arch/models"
	"context"
	"log"
)

func (a *authUseCase) SignUp(ctx context.Context, user *models.User) (*models.User, error) {
	log.Println("Function SignUp() is not implemented yet")
	return nil, nil
}
