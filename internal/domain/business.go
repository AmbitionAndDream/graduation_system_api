package domain

type ResponseBusinessList struct {
	Total  int64              `json:"total"`
	Limit  int                `json:"limit"`
	Offset int                `json:"offset"`
	Bus    []ResponseBusiness `json:"business"`
}
type ResponseBusiness struct {
	BusinessID   int    `json:"business_id"`
	BusinessName string `json:"business_name"`
}

type RequestDeleteBusinessId struct {
	BusinessId []int `json:"business_id" binding:"required"`
}

type RequestBusinessName struct {
	Name string `json:"name" binding:"required"`
}
