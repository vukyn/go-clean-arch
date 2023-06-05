package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	*gorm.Model
	Content   string
	CreatedBy string
	User      User `gorm:"foreignKey:CreatedBy"`
}
