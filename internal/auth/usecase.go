package auth

import (
	"boilerplate-clean-arch/internal/models"
	"context"

	"github.com/google/uuid"
)

type UseCase interface {
	SignUp(ctx context.Context, user *models.User) (*models.User, error)
	SignIn(ctx context.Context, email, password string) (*models.UserWithToken, error)
	GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error)
	// ParseToken(ctx context.Context, accessToken string) (string, error)
}
