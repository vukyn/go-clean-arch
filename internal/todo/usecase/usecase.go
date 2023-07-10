package usecase

import (
	"context"
	"go-clean-arch/internal/constants"
	commonModel "go-clean-arch/internal/models"
	"go-clean-arch/internal/todo/entity"
	"go-clean-arch/internal/todo/models"
	"go-clean-arch/internal/todo/repository"
	"go-clean-arch/pkg/utils"

	"github.com/labstack/gommon/log"
)

type usecase struct {
	repo repository.IRepository
}

// Constructor
func NewUseCase(repo repository.IRepository) IUseCase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) GetById(ctx context.Context, id int) (*models.TodoResponse, error) {
	record, err := u.repo.GetById(ctx, id)
	if err != nil {
		log.Errorf("usecase.repo.GetById: %v", err)
		return nil, utils.NewError(constants.STATUS_CODE_INTERNAL_SERVER, "Error when get todo")
	}
	return record.Export(), nil
}

func (u *usecase) GetList(ctx context.Context, params *models.RequestList) ([]*models.TodoResponse, error) {
	queries := params.ToMap()
	records, err := u.repo.GetList(ctx, queries)
	if err != nil {
		log.Errorf("usecase.repo.GetList: %v", err)
		return nil, utils.NewError(constants.STATUS_CODE_INTERNAL_SERVER, "Error when get list todo")
	}
	return (&entity.Todo{}).ExportList(records), nil
}

func (u *usecase) GetListPaging(ctx context.Context, params *models.RequestList) (*models.TodoListPaging, error) {
	queries := params.ToMap()
	records, err := u.repo.GetListPaging(ctx, queries)
	if err != nil {
		log.Errorf("usecase.repo.GetListPaging: %v", err)
		return nil, utils.NewError(constants.STATUS_CODE_INTERNAL_SERVER, "Error when get list todo")
	}
	count, err := u.repo.Count(ctx, queries)
	if err != nil {
		log.Errorf("usecase.repo.Count: %v", err)
		return nil, utils.NewError(constants.STATUS_CODE_INTERNAL_SERVER, "Error when get list todo")
	}

	return &models.TodoListPaging{
		ListPaging: commonModel.ListPaging{
			Page:  params.Page,
			Size:  params.Size,
			Total: count,
		},
		Records: (&entity.Todo{}).ExportList(records),
	}, nil
}

func (u *usecase) GetOne(ctx context.Context, params *models.RequestList) (*models.TodoResponse, error) {
	queries := params.ToMap()
	record, err := u.repo.GetOne(ctx, queries)
	if err != nil {
		log.Errorf("usecase.repo.GetOne: %v", err)
		return nil, utils.NewError(constants.STATUS_CODE_INTERNAL_SERVER, "Error when get todo")
	}
	return record.Export(), nil
}

func (u *usecase) Create(ctx context.Context, userId int, params *models.SaveRequest) (*models.TodoResponse, error) {
	obj := &entity.Todo{}
	obj.ParseForCreate(params, userId)
	res, err := u.repo.Create(ctx, obj)
	if err != nil {
		log.Errorf("usecase.repo.Create: %v", err)
		return nil, utils.NewError(constants.STATUS_CODE_INTERNAL_SERVER, "Error when create todo")
	}
	return res.Export(), nil
}

func (u *usecase) CreateMany(ctx context.Context, userId int, params []*models.SaveRequest) (int, error) {
	objs := (&entity.Todo{}).ParseForCreateMany(params, userId)
	res, err := u.repo.CreateMany(ctx, objs)
	if err != nil {
		log.Errorf("usecase.repo.Create: %v", err)
		return 0, utils.NewError(constants.STATUS_CODE_INTERNAL_SERVER, "Error when create todo")
	}
	return res, nil
}

func (u *usecase) Update(ctx context.Context, userId int, params *models.SaveRequest) (*models.TodoResponse, error) {
	obj := &entity.Todo{}
	obj.ParseForUpdate(params, userId)
	res, err := u.repo.Update(ctx, obj)
	if err != nil {
		log.Errorf("usecase.repo.Update: %v", err)
		return nil, utils.NewError(constants.STATUS_CODE_INTERNAL_SERVER, "Error when update todo")
	}
	return res.Export(), nil
}

func (u *usecase) UpdateMany(ctx context.Context, userId int, params []*models.SaveRequest) (int, error) {
	objs := (&entity.Todo{}).ParseForUpdateMany(params, userId)
	res, err := u.repo.UpdateMany(ctx, objs)
	if err != nil {
		log.Errorf("usecase.repo.Update: %v", err)
		return 0, utils.NewError(constants.STATUS_CODE_INTERNAL_SERVER, "Error when update todo")
	}
	return res, nil
}
