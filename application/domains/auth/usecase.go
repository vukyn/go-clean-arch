package auth

import (
	"boilerplate-clean-arch/models"
	"context"
)

type UseCase interface {
	SignUp(ctx context.Context, user *models.User) (*models.User, error)
	// SignIn(ctx context.Context, username, password string) (string, error)
	// ParseToken(ctx context.Context, accessToken string) (string, error)
}
