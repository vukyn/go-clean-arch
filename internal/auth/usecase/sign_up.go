package usecase

import (
	"boilerplate-clean-arch/internal/constants"
	"boilerplate-clean-arch/internal/models"
	"boilerplate-clean-arch/pkg/utils"
	"context"

	"github.com/labstack/gommon/log"
)

//	func (a *authUseCase) SignUp(ctx context.Context, user *models.User) (*models.User, error) {
//		log.Println("Function SignUp() is not implemented yet")
//		return nil, nil
//	}
func (a *authUseCase) SignUp(ctx context.Context, user *models.User) (*models.User, error) {
	log.SetPrefix("[SignUp]")
	log.Infof("Sign up new user with params: {FirstName: %s, LastName: %s, Email: %s}", user.FirstName, user.LastName, user.Email)

	// validation

	// check if user already exists
	_, err := a.userRepo.FindByEmail(ctx, user)
	if err == nil {
		log.Errorf("User already exist with email: %v", err)
		return nil, utils.NewError(constants.ERROR_CODE_BAD_REQUEST, constants.ERROR_MESSAGE_EMAIL_ALREADY_EXISTS)
	}

	if user.Gender != "Male" && user.Gender != "Female" {
		log.Errorf("Invalid gender type: %s", user.Gender)
		return nil, utils.NewError(constants.ERROR_CODE_BAD_REQUEST, constants.ERROR_MESSAGE_INVALID_GENDER_TYPE)
	}
	// end validation

	// create new user
	user, err = a.userRepo.SignUp(ctx, user.Parse())
	if err != nil {
		log.Errorf("Error while creating new user: %s", err)
		return nil, err
	}

	return user, nil
}
