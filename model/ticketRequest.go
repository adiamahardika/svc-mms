package model

import (
	"encoding/json"
)

type GetTicketRequest struct {
	PageNo 			json.Number `json:"page_no"`
	PageSize 		json.Number `json:"page_size"`
	SortBy 			string `json:"sort_by"`
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
	Kategori		string `json:"kategori"`
	Email			string `json:"email"`
	AssignedTo		string `json:"assigned_to"`
	AssignedToTeam	string `json:"assigned_to_team"`
	TicketCode		string `json:"ticket_code"`
}

type AssignTicketToMemberRequest struct {
	Id		json.Number	`json:"id"`
	UserId	json.Number	`json:"user_id"`
}