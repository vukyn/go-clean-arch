package usecase

import (
	"boilerplate-clean-arch/internal/constants"
	"boilerplate-clean-arch/internal/models"
	"boilerplate-clean-arch/pkg/utils"
	"context"

	"github.com/labstack/gommon/log"
)

//	func (a *authUseCase) SignIn(ctx context.Context, username, password string) (string, error) {
//		log.Println("Function SignIn() is not implemented yet")
//		return nil, nil
//	}
func (a *authUseCase) SignIn(ctx context.Context, email, password string) (*models.UserWithToken, error) {
	log.SetPrefix("[SignIn]")
	log.Infof("Sign in with user {Email: %s}", email)

	// validation

	// check if user already exists
	foundUser, err := a.userRepo.FindByEmail(ctx, email)
	if err != nil {
		log.Errorf("Error while finding user by email: %s", err)
		return nil, utils.NewError(constants.ERROR_CODE_BAD_REQUEST, constants.ERROR_MESSAGE_INTERNAL_SERVER_ERROR)
	}
	if foundUser == nil {
		log.Errorf("User not found with email: %v", err)
		return nil, utils.NewError(constants.ERROR_CODE_BAD_REQUEST, constants.ERROR_MESSAGE_USER_NOT_FOUND)
	}

	// check if password is correct
	if err = foundUser.ComparePasswords(password); err != nil {
		log.Errorf("Compare password failed: %v", err)
		return nil, utils.NewError(constants.ERROR_CODE_UNAUTHORIZED, constants.ERROR_MESSAGE_INVALID_EMAIL_OR_PASSWORD)
	}
	// end validation

	// generate token
	token, err := utils.GenerateJWTToken(foundUser, a.cfg)
	if err != nil {
		log.Errorf("Cannot generate token: %v", err)
		return nil, utils.NewError(constants.ERROR_CODE_INTERNAL_SERVER, constants.ERROR_MESSAGE_INTERNAL_SERVER_ERROR)
	}
	foundUser.SanitizePassword()

	// create session
	_, err = a.sessRepo.CreateSession(ctx, "api:sign-in", &models.Session{
		UserId: foundUser.UserId,
	}, a.cfg.Auth.Expire)
	if err != nil {
		log.Errorf("authUC.sessRepo.CreateSession: %v", err)
		return nil, utils.NewError(constants.ERROR_CODE_INTERNAL_SERVER, constants.ERROR_MESSAGE_INTERNAL_SERVER_ERROR)
	}

	return &models.UserWithToken{
		User:  foundUser,
		Token: token,
	}, nil
}
