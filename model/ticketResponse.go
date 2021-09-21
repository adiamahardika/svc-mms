package model

type CountTicketByStatusResponse struct {
	Status string `json:"status"`
	Total  int    `json:"total"`
}