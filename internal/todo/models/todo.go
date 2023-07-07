package models

import (
	commonModel "boilerplate-clean-arch/internal/models"
	"time"
)

type RequestList struct {
	commonModel.RequestPaging
	FromDate  int64
	ToDate    int64
	CreatedBy int64
}

func (r *RequestList) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"from_date":  r.FromDate,
		"to_date":    r.ToDate,
		"created_by": r.CreatedBy,
		"page":       r.Page,
		"size":       r.Size,
		"sort_by":    r.SortBy,
		"order_by":   r.OrderBy,
	}
}

type TodoResponse struct {
	Id        int
	Content   string
	CreatedBy int
	CreatedAt time.Time
	UpdateBy  int
	UpdatedAt time.Time
}

type SaveRequest struct {
	Id      int64
	Content string
}

type ListPaging struct {
	commonModel.ListPaging
	Records []*TodoResponse
}