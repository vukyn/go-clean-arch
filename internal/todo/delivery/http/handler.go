package http

import (
	"boilerplate-clean-arch/internal/todo"
	"boilerplate-clean-arch/config"
)

type todoHandlers struct {
	cfg    *config.Config
	todoUC todo.UseCase
}

func NewTodoHandlers(cfg *config.Config, todoUC todo.UseCase) todo.Handlers {
	return &todoHandlers{
		cfg:    cfg,
		todoUC: todoUC,
	}
}