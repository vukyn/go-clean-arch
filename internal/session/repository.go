package session

import (
	"boilerplate-clean-arch/internal/models"
	"context"
)

// Session repository
type SessRepository interface {
	CreateSession(ctx context.Context, prefix string, session *models.Session, expire int) (string, error)
	GetSessionByID(ctx context.Context, sessionID string) (*models.Session, error)
	DeleteByID(ctx context.Context, sessionID string) error
}
