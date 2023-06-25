package repository

import (
	"boilerplate-clean-arch/internal/models"
	"context"
)

func (t *todoRepo) Create(ctx context.Context, todo *models.Todo) (int64, error) {
	result := t.db.Create(todo)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
