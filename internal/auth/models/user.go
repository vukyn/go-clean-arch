package models

import (
	commonModel "boilerplate-clean-arch/internal/models"
	"time"

	"github.com/google/uuid"
)

type RequestList struct {
	commonModel.RequestPaging
	Email string
}

func (r *RequestList) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"email":    r.Email,
		"page":     r.Page,
		"size":     r.Size,
		"sort_by":  r.SortBy,
		"order_by": r.OrderBy,
	}
}

type UserResponse struct {
	Id          int
	UserId      uuid.UUID
	FirstName   string
	LastName    string
	Email       string
	Role        string
	About       string
	Avatar      string
	PhoneNumber string
	Address     string
	City        string
	Country     string
	Gender      string
	Birthday    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	LoginDate   time.Time
}

type SaveRequest struct {
	Id        int64
	FirstName string
	LastName  string
	Email     string
	Password  string
	Gender    string
	City      string
	Country   string
	Birthday  time.Time
}

type LoginRequest struct {
	Email    string
	Password string
}

// User sign in response
type UserWithToken struct {
	User  *UserResponse `json:"user,omitempty"`
	Token string        `json:"token"`
}
