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
	Id          int        `json:"id,omitempty"`
	UserId      uuid.UUID  `json:"user_id,omitempty"`
	FirstName   string     `json:"first_name,omitempty"`
	LastName    string     `json:"last_name,omitempty"`
	Email       string     `json:"email,omitempty"`
	Role        string     `json:"role,omitempty"`
	About       string     `json:"about,omitempty"`
	Avatar      string     `json:"avatar,omitempty"`
	PhoneNumber string     `json:"phone_number,omitempty"`
	Address     string     `json:"address,omitempty"`
	City        string     `json:"city,omitempty"`
	Country     string     `json:"country,omitempty"`
	Gender      string     `json:"gender,omitempty"`
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
