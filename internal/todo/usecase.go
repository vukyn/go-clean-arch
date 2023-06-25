package todo

import (
	"boilerplate-clean-arch/internal/models"
	"context"
)

type UseCase interface {
	Create(ctx context.Context, todo *models.Todo) (int64, error)
	// Update(ctx context.Context, todo *models.todo) (*models.todo, error)
	// GettodoByID(ctx context.Context, todoID uuid.UUID) (*models.todoBase, error)
	// Delete(ctx context.Context, todoID uuid.UUID) error
	// GetTodos(ctx context.Context, pq *utils.PaginationQuery) (*models.todoList, error)
	// SearchByTitle(ctx context.Context, title string, query *utils.PaginationQuery) (*models.todoList, error)
}
