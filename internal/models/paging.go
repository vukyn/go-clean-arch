package models

type RequestPaging struct {
	Page    int
	Size    int
	SortBy  string
	OrderBy string
}

type ListPaging struct {
	Page    int
	Size    int
	Total   int
	Records interface{}
}
