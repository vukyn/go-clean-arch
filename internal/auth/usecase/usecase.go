package usecase

import (
	"boilerplate-clean-arch/config"
	"boilerplate-clean-arch/internal/auth"
	"fmt"
)

type authUseCase struct {
	cfg       *config.Config
	authRepo  auth.AuthRepository
	redisRepo auth.RedisRepository
}

// Constructor
func NewAuthUseCase(cfg *config.Config, authRepo auth.AuthRepository, redisRepo auth.RedisRepository) auth.UseCase {
	return &authUseCase{
		cfg:       cfg,
		authRepo:  authRepo,
		redisRepo: redisRepo,
	}
}

const (
	basePrefix    = "api-auth:"
	cacheDuration = 3600
)

func (u *authUseCase) GenerateUserKey(userID string) string {
	return fmt.Sprintf("%s: %s", basePrefix, userID)
}
