package model

import (
	"encoding/json"
)

type GetTicketRequest struct {
	PageNo 			json.Number `json:"page_no" binding:"required"`
	PageSize 		json.Number `json:"page_size" binding:"required"`
	SortBy 			string `json:"sort_by"`
	Search 			string `json:"search"`
	Status 			string `json:"status"`
	Priority 		string `json:"priority"`
	AssignedTo		string `json:"assigned_to"`
	AssignedToTeam	string `json:"assigned_to_team"`
	UsernamePembuat string `json:"username_pembuat"`
	StartDate 		string `json:"start_date"`
	EndDate 		string `json:"end_date"`
}