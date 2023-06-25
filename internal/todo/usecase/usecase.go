package usecase

import (
	"boilerplate-clean-arch/internal/todo"
)

type todoUseCase struct {
	todoRepo todo.Repository
}

// Constructor
func NewTodoUseCase(todoRepo todo.Repository) todo.UseCase {
	return &todoUseCase{
		todoRepo: todoRepo,
	}
}
