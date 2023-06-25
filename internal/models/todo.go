package models

import (
	"time"
)

type Todo struct {
	Id        int       `gorm:"primarykey;column:id" json:"id" redis:"id"`
	Content   string    `gorm:"column:content" json:"content,omitempty" redis:"content" validate:"required,lte=1024"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at,omitempty" redis:"created_at" validate:"required"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at,omitempty" redis:"updated_at"`
	CreatedBy int       `gorm:"column:created_by" json:"created_by,omitempty" redis:"created_by" validate:"required"`
	UpdateBy  int       `gorm:"column:update_by" json:"update_by,omitempty" redis:"update_by"`
}
