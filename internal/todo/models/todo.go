package models

import (
	commonModel "boilerplate-clean-arch/internal/models"
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
	Id        int    `json:"id"`
	Content   string `json:"content,omitempty"`
	CreatedBy int    `json:"created_by,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdateBy  int    `json:"update_by,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type SaveRequest struct {
	Id      int64  `json:"id"`
	Content string `json:"content"`
}

type ListPaging struct {
	commonModel.ListPaging
	Records []*TodoResponse
}
