package auth

import (
	"boilerplate-clean-arch/internal/models"
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	SignUp(ctx context.Context, user *models.User) (int64, error)
	// Update(ctx context.Context, user *models.User) (*models.User, error)
	// Delete(ctx context.Context, userID uuid.UUID) error
	GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error)
	// FindByName(ctx context.Context, name string, query *utils.PaginationQuery) (*models.UsersList, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	// GetUsers(ctx context.Context, pq *utils.PaginationQuery) (*models.UsersList, error)
}
