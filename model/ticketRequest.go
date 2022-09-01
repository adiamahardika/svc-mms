package model

import (
	"encoding/json"
	"time"
)

type GetTicketRequest struct {
	AssignedTo      string   `json:"assigned_to"`
	Category        []string `json:"category"`
	PageNo          int      `json:"page_no"`
	PageSize        int      `json:"page_size"`
	StartIndex      int      `json:"start_index"`
	Priority        string   `json:"priority"`
	Search          string   `json:"search"`
	Status          string   `json:"status"`
	AssignedToTeam  string   `json:"assigned_to_team"`
	AreaCode        []string `json:"area_code"`
	Regional        []string `json:"regional"`
	GrapariId       []string `json:"grapari_id"`
	UsernamePembuat string   `json:"username_pembuat"`
	StartDate       string   `json:"start_date" binding:"required"`
	EndDate         string   `json:"end_date" binding:"required"`
	VisitStatus     []string `json:"visit_status"`
}

type CreateTicketRequest struct {
	Judul          string `json:"judul"`
	UserPembuat    string `json:"user_pembuat"`
	Prioritas      string `json:"prioritas"`
	Status         string `json:"status"`
	Isi            string `json:"isi"`
	Lokasi         string `json:"lokasi"`
	TerminalId     string `json:"terminal_id"`
	Category       string `json:"category"`
	Email          string `json:"email"`
	AssignedTo     string `json:"assigned_to"`
	AssignedToTeam string `json:"assigned_to_team"`
	TicketCode     string `json:"ticket_code"`
	NoSPM          string `json:"no_spm"`
	NoReqSPM       string `json:"no_req_spm"`
	AreaCode       string `json:"area_code"`
	Regional       string `json:"regional"`
	GrapariId      string `json:"grapari_id"`
	SubCategory    string `json:"sub_category"`
}

type AssignTicketRequest struct {
	TicketCode    string      `json:"ticket_code" binding:"required"`
	AssigningBy   json.Number `json:"assigning_by"`
	AssigningTime time.Time   `json:"assigning_time"`
	UserId        json.Number `json:"user_id" binding:"required"`
	TeamId        json.Number `json:"team_id" binding:"required"`
	UpdateAt      time.Time   `json:"update_at"`
}

type UpdateTicketStatusRequest struct {
	TicketCode string    `json:"ticket_code" binding:"required"`
	Status     string    `json:"status" binding:"required"`
	UpdateAt   time.Time `json:"update_at"`
}

type CountTicketByStatusRequest struct {
	AreaCode       []string `json:"area_code"`
	Regional       []string `json:"regional"`
	GrapariId      []string `json:"grapari_id"`
	AssignedTo     string   `json:"assigned_to"`
	AssignedToTeam string   `json:"assigned_to_team"`
	StartDate      string   `json:"start_date" binding:"required"`
	EndDate        string   `json:"end_date" binding:"required"`
}

type GetEmailHistoryRequest struct {
	Search string `json:"search"`
}

type UpdateVisitStatusRequest struct {
	TicketCode  string    `json:"ticket_code" binding:"required"`
	VisitStatus string    `json:"visit_status" binding:"required"`
	UpdateAt    time.Time `json:"update_at"`
}
