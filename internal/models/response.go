package models

type Response struct {
	Pagination Pagination `json:"pagination"`
	Services   []Service  `json:"services"`
}
