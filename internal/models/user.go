package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id          int       `gorm:"primarykey;column:id" json:"id" redis:"id"`
	UserId      uuid.UUID `gorm:"column:user_id" json:"user_id" redis:"user_id" validate:"omitempty"`
	FirstName   string    `gorm:"column:first_name" json:"first_name" redis:"first_name" validate:"required,lte=30"`
	LastName    string    `gorm:"column:last_name" json:"last_name" redis:"last_name" validate:"required,lte=30"`
	Email       string    `gorm:"column:email" json:"email,omitempty" redis:"email" validate:"omitempty,lte=60,email"`
	Password    string    `gorm:"column:password" json:"password,omitempty" redis:"password" validate:"omitempty,required,gte=6"`
	Role        string    `gorm:"column:role" json:"role,omitempty" redis:"role" validate:"omitempty,lte=10"`
	About       string    `gorm:"column:about" json:"about,omitempty" redis:"about" validate:"omitempty,lte=1024"`
	Avatar      string    `gorm:"column:avatar" json:"avatar,omitempty" redis:"avatar" validate:"omitempty,lte=512,url"`
	PhoneNumber string    `gorm:"column:phone_number" json:"phone_number,omitempty" redis:"phone_number" validate:"omitempty,lte=20"`
	Address     string    `gorm:"column:address" json:"address,omitempty" redis:"address" validate:"omitempty,lte=250"`
	City        string    `gorm:"column:city" json:"city,omitempty" redis:"city" validate:"omitempty,lte=24"`
	Country     string    `gorm:"column:country" json:"country,omitempty" redis:"country" validate:"omitempty,lte=24"`
	Gender      string    `gorm:"column:gender" json:"gender,omitempty" redis:"gender" validate:"omitempty,lte=10"`
	Birthday    time.Time `gorm:"column:birthday" json:"birthday,omitempty" redis:"birthday" validate:"omitempty,lte=10"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at,omitempty" redis:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at,omitempty" redis:"updated_at"`
	LoginDate   time.Time `gorm:"column:login_date" json:"login_date" redis:"login_date"`
}

// Hash user password with bcrypt
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Compare user password and payload
func (u *User) ComparePasswords(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

// Sanitize user password
func (u *User) SanitizePassword() {
	u.Password = ""
}
