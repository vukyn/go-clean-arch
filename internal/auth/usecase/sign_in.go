package usecase

import (
	"boilerplate-clean-arch/internal/constants"
	"boilerplate-clean-arch/internal/models"
	"boilerplate-clean-arch/pkg/httpErrors"
	"boilerplate-clean-arch/pkg/utils"
	"context"

	"github.com/labstack/gommon/log"
)

//	func (a *authUseCase) SignIn(ctx context.Context, username, password string) (string, error) {
//		log.Println("Function SignIn() is not implemented yet")
//		return nil, nil
//	}
func (a *authUseCase) SignIn(ctx context.Context, email, password string) (string, error) {
	log.SetPrefix("[SignIn]")
	log.Infof("Sign in with user {Email: %s}", email)

	// validation

	// check if user already exists
	user, err := a.userRepo.FindByEmail(ctx, &models.User{Email: email})
	if err != nil {
		log.Errorf("User not found with email: %v", err)
		return "", utils.NewError(constants.ERROR_CODE_BAD_REQUEST, constants.ERROR_MESSAGE_USER_NOT_FOUND)
	}

	// check if password is correct
	if err = user.ComparePasswords(password); err != nil {
		return "", utils.NewError(constants.ERROR_CODE_UNAUTHORIZED, constants.ERROR_MESSAGE_INVALID_EMAIL_OR_PASSWORD)
	}
	// end validation

	// generate token
	token, err := utils.GenerateJWTToken(user, a.cfg)
	if err != nil {
		log.Errorf("Cannot generate token: %v", err)
		return "", utils.NewError(constants.ERROR_CODE_INTERNAL_SERVER, constants.ERROR_MESSAGE_INTERNAL_SERVER_ERROR)
	}

	return token, nil
}
