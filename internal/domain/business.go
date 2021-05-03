package domain

type ResponseBusinessList struct {
	Total  int        `json:"total"`
	Limit  int        `json:"limit"`
	Offset int        `json:"offset"`
	Bus    []ResponseBusiness `json:"business"`
}
type ResponseBusiness struct {
	BusinessID   int    `json:"business_id"`
	BusinessName string `json:"business_name"`
}
