package usecase

import (
	"boilerplate-clean-arch/internal/todo/models"
	"context"
)

type IUseCase interface {
	Create(ctx context.Context, userId int, params *models.SaveRequest) (*models.Response, error)
	CreateMany(ctx context.Context, userId int, params []*models.SaveRequest) (int, error)
	Update(ctx context.Context, userId int, params *models.SaveRequest) (*models.Response, error)
	UpdateMany(ctx context.Context, userId int, params []*models.SaveRequest) (int, error)
	GetById(ctx context.Context, id int) (*models.Response, error)
	GetList(ctx context.Context, params *models.RequestList) ([]*models.Response, error)
	GetListPaging(ctx context.Context, params *models.RequestList) (*models.ListPaging, error)
	GetOne(ctx context.Context, params *models.RequestList) (*models.Response, error)
}
