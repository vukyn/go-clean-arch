package usecase

import (
	"boilerplate-clean-arch/internal/todo/models"
	"context"
)

type IUseCase interface {
	Create(ctx context.Context, userId int, params *models.SaveRequest) (*models.TodoResponse, error)
	CreateMany(ctx context.Context, userId int, params []*models.SaveRequest) (int, error)
	Update(ctx context.Context, userId int, params *models.SaveRequest) (*models.TodoResponse, error)
	UpdateMany(ctx context.Context, userId int, params []*models.SaveRequest) (int, error)
	GetById(ctx context.Context, id int) (*models.TodoResponse, error)
	GetList(ctx context.Context, params *models.RequestList) ([]*models.TodoResponse, error)
	GetListPaging(ctx context.Context, params *models.RequestList) (*models.TodoListPaging, error)
	GetOne(ctx context.Context, params *models.RequestList) (*models.TodoResponse, error)
}
