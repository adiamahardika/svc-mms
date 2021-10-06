package model

import (
	"encoding/json"
	"time"
)

type GetTicketRequest struct {
	Search 			string `json:"search"`
	Status 			string `json:"status"`
	Priority 		string `json:"priority"`
	AssignedTo		string `json:"assigned_to"`
	AssignedToTeam	string `json:"assigned_to_team"`
	UsernamePembuat string `json:"username_pembuat"`
	StartDate 		string `json:"start_date" binding:"required"`
	EndDate 		string `json:"end_date" binding:"required"`
}

type CreateTicketRequest struct {
	Judul 			string `json:"judul"`
	UserPembuat 	string `json:"user_pembuat"`
	Prioritas 		string `json:"prioritas"`
	TotalWaktu		string `json:"total_waktu"`
	Status			string `json:"status"`
	Isi				string `json:"isi"`
	Lokasi			string `json:"lokasi"`
	TerminalId		string `json:"terminal_id"`
	Category		string `json:"category"`
	Email			string `json:"email"`
	AssignedTo		string `json:"assigned_to"`
	AssignedToTeam	string `json:"assigned_to_team"`
	TicketCode		string `json:"ticket_code"`
}

type AssignTicketRequest struct {
	TicketCode	string	`json:"ticket_code" binding:"required"`
	UserId		json.Number	`json:"user_id" binding:"required"`
	TeamId		json.Number	`json:"team_id" binding:"required"`
	UpdateAt 	time.Time `json:"update_at"`
}

type UpdateTicketStatusRequest struct {
	TicketCode	string	`json:"ticket_code" binding:"required"`
	Status		string	`json:"status" binding:"required"`
	UpdateAt 	time.Time `json:"update_at"`
}