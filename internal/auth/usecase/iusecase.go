package usecase

import (
	"boilerplate-clean-arch/internal/auth/models"
	"context"
)

type IUseCase interface {
	Register(ctx context.Context, params *models.SaveRequest) (*models.UserResponse, error)
	Login(ctx context.Context, params *models.LoginRequest) (*models.UserWithToken, error)
	// GetByEmail(ctx context.Context, email string) (*models.Response, error)
}