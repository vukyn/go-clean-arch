package usecase

import (
	"context"
	"go-clean-arch/internal/auth/models"
)

type IUseCase interface {
	Register(ctx context.Context, params *models.SaveRequest) (*models.UserResponse, error)
	Login(ctx context.Context, params *models.LoginRequest) (*models.UserWithToken, error)
	// GetByEmail(ctx context.Context, email string) (*models.Response, error)
}
