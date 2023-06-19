package usecase

import (
	"boilerplate-clean-arch/internal/models"
	"context"
	"github.com/labstack/gommon/log"
)

// func (a *authUseCase) SignUp(ctx context.Context, user *models.User) (*models.User, error) {
// 	log.Println("Function SignUp() is not implemented yet")
// 	return nil, nil
// }
func (a *authUseCase) SignUp(ctx context.Context, user *models.User) (*models.User, error) {
	log.Infof("Sign up new user with params: %#v", user)

	// check if user already exists
	_, err := a.userRepo.FindByEmail(ctx, user)
	if err != nil {
		log.Errorf("Error while getting user by email: %s", err)
		return nil, err
	}

	// create new user
	user, err = a.userRepo.SignUp(ctx, user)
	if err != nil {
		log.Errorf("Error while creating new user: %s", err)
		return nil, err
	}

	return user, nil
}
