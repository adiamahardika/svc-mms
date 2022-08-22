package model

import "time"

type CreatePreventiveRequest struct {
	VisitDate      string `json:"visit_date"`
	Location       string `json:"location"`
	TerminalId     string `json:"terminal_id"`
	AssignedTo     string `json:"assigned_to"`
	AssignedToTeam string `json:"assigned_to_team"`
	Status         string `json:"status"`
	CreatedBy      string `json:"created_by"`
}

type GetPreventiveRequest struct {
	Search         string   `json:"search"`
	Status         []string `json:"status"`
	AssignedTo     string   `json:"assigned_to"`
	AssignedToTeam string   `json:"assigned_to_team"`
	StartDate      string   `json:"start_date" binding:"required"`
	EndDate        string   `json:"end_date" binding:"required"`
	AreaCode       []string `json:"area_code"`
	Regional       []string `json:"regional"`
	GrapariId      []string `json:"grapari_id"`
	PageNo         int      `json:"page_no"`
	PageSize       int      `json:"page_size"`
	StartIndex     int      `json:"start_index"`
}

type UpdatePreventiveRequest struct {
	PrevCode       string    `json:"prev_code"`
	VisitDate      string    `json:"visit_date"`
	Location       string    `json:"location"`
	TerminalId     string    `json:"terminal_id"`
	AssignedTo     string    `json:"assigned_to"`
	AssignedToTeam string    `json:"assigned_to_team"`
	Status         string    `json:"status"`
	UpdatedBy      string    `json:"updated_by"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CountPreventiveByStatusRequest struct {
	AssignedTo     string `json:"assigned_to"`
	AssignedToTeam string `json:"assigned_to_team"`
	StartDate      string `json:"start_date" binding:"required"`
	EndDate        string `json:"end_date" binding:"required"`
}
