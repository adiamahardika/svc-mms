package model

import "time"

type CreatePreventiveRequest struct {
	VisitDate  string `json:"visit_date"`
	TerminalId string `json:"terminal_id"`
	AssignedTo string `json:"assigned_to"`
	Status     string `json:"status"`
	CreatedBy  string `json:"created_by"`
}

type GetPreventiveRequest struct {
	Search     string `json:"search"`
	Status     string `json:"status"`
	AssignedTo string `json:"assigned_to"`
	StartDate  string `json:"start_date" binding:"required"`
	EndDate    string `json:"end_date" binding:"required"`
}

type UpdatePreventiveRequest struct {
	PrevCode   string    `json:"prev_code"`
	VisitDate  string    `json:"visit_date"`
	TerminalId string    `json:"terminal_id"`
	AssignedTo string    `json:"assigned_to"`
	Status     string    `json:"status"`
	UpdatedBy  string    `json:"updated_by"`
	UpdatedAt  time.Time `json:"updated_at"`
}