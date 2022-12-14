package model

import (
	"svc-monitoring-maintenance/entity"
	"time"
)

type CountTicketByStatusResponse struct {
	Status string `json:"status"`
	Total  int    `json:"total"`
}

type GetTicketResponse struct {
	Id               int                `json:"id" gorm:"primaryKey"`
	Judul            string             `json:"judul"`
	UsernamePembuat  string             `json:"username_pembuat"`
	UsernamePembalas string             `json:"username_pembalas"`
	Prioritas        string             `json:"prioritas"`
	TglDibuat        time.Time          `json:"tgl_dibuat"`
	TglDiperbarui    time.Time          `json:"tgl_diperbarui"`
	TotalWaktu       string             `json:"total_waktu"`
	Status           string             `json:"status"`
	TicketCode       string             `json:"ticket_code"`
	Category         string             `json:"category"`
	CategoryName     string             `json:"category_name"`
	SubCategory      string             `json:"sub_category"`
	AreaCode         string             `json:"area_code"`
	AreaName         string             `json:"area_name"`
	Regional         string             `json:"regional"`
	GrapariId        string             `json:"grapari_id"`
	GrapariName      string             `json:"grapari_name"`
	Lokasi           string             `json:"lokasi"`
	TerminalId       string             `json:"terminal_id"`
	Email            string             `json:"email"`
	AssignedTo       string             `json:"assigned_to"`
	UserName         string             `json:"user_name"`
	AssignedToTeam   string             `json:"assigned_to_team"`
	TeamName         string             `json:"team_name"`
	TicketIsi        []entity.TicketIsi `json:"ticket_isi" gorm:"foreignKey:TicketCode"`
	NoSPM            string             `json:"no_spm"`
	NoReqSPM         string             `json:"no_req_spm"`
	UserPembuat      string             `json:"user_pembuat"`
	Assignee         string             `json:"assignee"`
}

type GetEmailHistoryResponse struct {
	Email string `json:"email"`
}

type GetTicketActivityResponse struct {
	Date         time.Time `json:"date"`
	Unassigned   int       `json:"unassigned"`
	Investigated int       `json:"investigated"`
	Close        int       `json:"close"`
}
