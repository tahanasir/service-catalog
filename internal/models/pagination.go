package models

type Pagination struct {
	TotalServices int `json:"totalServices"`
	TotalPages    int `json:"totalPages"`
	CurrentPage   int `json:"currentPage"`
	NextPage      int `json:"nextPage,omitempty"`
	PrevPage      int `json:"prevPage,omitempty"`
}
