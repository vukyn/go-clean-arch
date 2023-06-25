package repository

import (
	"boilerplate-clean-arch/internal/todo"

	"gorm.io/gorm"
)

// Todo repository
type todoRepo struct {
	db *gorm.DB
}

// Constructor
func NewTodoRepository(db *gorm.DB) todo.Repository {
	return &todoRepo{
		db: db,
	}
}