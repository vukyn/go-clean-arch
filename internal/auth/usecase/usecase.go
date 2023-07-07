package usecase

import (
	"boilerplate-clean-arch/config"
	"boilerplate-clean-arch/internal/auth/entity"
	"boilerplate-clean-arch/internal/auth/models"
	"boilerplate-clean-arch/internal/auth/repository"
	"boilerplate-clean-arch/internal/constants"
	"boilerplate-clean-arch/pkg/utils"
	"context"
	"fmt"

	"github.com/labstack/gommon/log"
)

type usecase struct {
	cfg       *config.Config
	repo      repository.IRepository
	redisRepo repository.IRedisRepository
}

// Constructor
func NewUseCase(cfg *config.Config, repo repository.IRepository, redisRepo repository.IRedisRepository) IUseCase {
	return &usecase{
		cfg:       cfg,
		repo:      repo,
		redisRepo: redisRepo,
	}
}

const (
	basePrefix    = "api-auth:"
	cacheDuration = 3600
)

func (u *usecase) GenerateUserKey(userId int) string {
	return fmt.Sprintf("%s: %d", basePrefix, userId)
}

func (u *usecase) Register(ctx context.Context, params *models.SaveRequest) (*models.UserResponse, error) {
	log.SetPrefix("[Register]")
	log.Infof("Register user with params: {FirstName: %s, LastName: %s, Email: %s}", params.FirstName, params.LastName, params.Email)

	// validation

	// check if user already exists
	foundUser, err := u.repo.GetOne(ctx, (&models.RequestList{Email: params.Email}).ToMap())
	if err != nil {
		log.Errorf("Error while finding user by email: %s", err)
		return nil, utils.NewError(constants.STATUS_CODE_BAD_REQUEST, constants.STATUS_MESSAGE_INTERNAL_SERVER_ERROR)
	}
	if foundUser.Id != 0 {
		log.Errorf("User already exist with email: %v", err)
		return nil, utils.NewError(constants.STATUS_CODE_BAD_REQUEST, constants.STATUS_MESSAGE_EMAIL_ALREADY_EXISTS)
	}

	if params.Gender != "Male" && params.Gender != "Female" {
		log.Errorf("Invalid gender type: %s", params.Gender)
		return nil, utils.NewError(constants.STATUS_CODE_BAD_REQUEST, constants.STATUS_MESSAGE_INVALID_GENDER_TYPE)
	}
	// end validation

	// create new user
	obj := &entity.User{}
	obj.HashPassword()
	obj.ParseFromSaveRequest(params)
	res, err := u.repo.Create(ctx, obj)
	if err != nil {
		log.Errorf("Error while creating new user: %s", err)
		return nil, err
	}
	res.SanitizePassword()
	return res.Export(), nil
}

func (u *usecase) Login(ctx context.Context, params *models.LoginRequest) (*models.UserWithToken, error) {
	log.SetPrefix("[Login]")
	log.Infof("Sign in with user {Email: %s}", params.Email)

	// validation

	// check if user already exists
	foundUser, err := u.repo.GetOne(ctx, (&models.RequestList{Email: params.Email}).ToMap())
	if err != nil {
		log.Errorf("Error while finding user by email: %s", err)
		return nil, utils.NewError(constants.STATUS_CODE_BAD_REQUEST, constants.STATUS_MESSAGE_INTERNAL_SERVER_ERROR)
	}
	if foundUser == nil {
		log.Errorf("User not found with email: %v", err)
		return nil, utils.NewError(constants.STATUS_CODE_BAD_REQUEST, constants.STATUS_MESSAGE_USER_NOT_FOUND)
	}

	// check if password is correct
	if err = utils.ComparePasswords(foundUser.Password, params.Password); err != nil {
		log.Errorf("Compare password failed: %v", err)
		return nil, utils.NewError(constants.STATUS_CODE_UNAUTHORIZED, constants.STATUS_MESSAGE_INVALID_EMAIL_OR_PASSWORD)
	}
	// end validation

	// generate token
	token, err := utils.GenerateJWTToken(foundUser.Export(), u.cfg)
	if err != nil {
		log.Errorf("Cannot generate token: %v", err)
		return nil, utils.NewError(constants.STATUS_CODE_INTERNAL_SERVER, constants.STATUS_MESSAGE_INTERNAL_SERVER_ERROR)
	}

	// save to cache
	if err = u.redisRepo.SetUser(ctx, u.GenerateUserKey(foundUser.Id), cacheDuration, foundUser); err != nil {
		log.Errorf("usecase.redisRepo.SetUser: %v", err)
		return nil, err
	}

	foundUser.SanitizePassword()

	return &models.UserWithToken{
		User:  foundUser.Export(),
		Token: token,
	}, nil
}

func (u *usecase) GetById(ctx context.Context, userId int) (*models.UserResponse, error) {

	cachedUser, err := u.redisRepo.GetById(ctx, u.GenerateUserKey(userId))
	if err != nil {
		log.Errorf("usecase.redisRepo.GetById: %v", err)
		return nil, err
	}
	if cachedUser != nil {
		return cachedUser.Export(), nil
	}

	user, err := u.repo.GetById(ctx, userId)
	if err != nil {
		return nil, err
	}

	if err = u.redisRepo.SetUser(ctx, u.GenerateUserKey(userId), cacheDuration, user); err != nil {
		log.Errorf("usecase.redisRepo.SetUser: %v", err)
		return nil, err
	}

	user.SanitizePassword()

	return user.Export(), nil
}
