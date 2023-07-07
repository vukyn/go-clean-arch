package entity

import (
	"boilerplate-clean-arch/internal/todo/models"
	"time"

	"github.com/jinzhu/copier"
)

type Todo struct {
	Id        int       `gorm:"primarykey;column:id" json:"id" redis:"id"`
	Content   string    `gorm:"column:content" json:"content,omitempty" redis:"content" validate:"required,lte=1024"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at,omitempty" redis:"created_at" validate:"required"`
	CreatedBy int       `gorm:"column:created_by" json:"created_by,omitempty" redis:"created_by" validate:"required"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;default:(-)" json:"updated_at,omitempty" redis:"updated_at"`
	UpdatedBy int       `gorm:"column:update_by;default:(-)" json:"update_by,omitempty" redis:"update_by"`
}

func (t *Todo) TableName() string {
	return "todos"
}

func (a *Todo) Export() *models.TodoResponse {
	obj := &models.TodoResponse{}
	copier.Copy(obj, a) //nolint
	return obj
}

func (a *Todo) ExportList(in []*Todo) []*models.TodoResponse {
	objs := make([]*models.TodoResponse, 0)
	for _, v := range in {
		objs = append(objs, v.Export())
	}
	return objs
}

func (a *Todo) ParseFromSaveRequest(req *models.SaveRequest) {
	copier.Copy(a, req) //nolint
}

func (a *Todo) ParseForCreate(req *models.SaveRequest, userId int) {
	a.ParseFromSaveRequest(req)
	a.CreatedBy = userId
}

func (a *Todo) ParseForCreateMany(reqs []*models.SaveRequest, userId int) []*Todo {
	objs := make([]*Todo, 0)
	for _, v := range reqs {
		obj := &Todo{}
		obj.ParseForCreate(v, userId)
		objs = append(objs, obj)
	}
	return objs
}

func (a *Todo) ParseForUpdate(req *models.SaveRequest, userId int) {
	a.ParseFromSaveRequest(req)
	a.UpdatedBy = userId
}

func (a *Todo) ParseForUpdateMany(reqs []*models.SaveRequest, userId int) []*Todo {
	objs := make([]*Todo, 0)
	for _, v := range reqs {
		obj := &Todo{}
		obj.ParseForUpdate(v, userId)
		objs = append(objs, obj)
	}
	return objs
}
