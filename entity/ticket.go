package entity

import "time"

type Ticket struct {
	Id              int       `json:"id" gorm:"primaryKey"`
	Judul           string    `json:"judul"`
	UsernamePembuat string    `json:"username_pembuat"`
	UpdatedBy       string    `json:"updated_by" form:"updatedBy"`
	Prioritas       string    `json:"prioritas"`
	TglDibuat       time.Time `json:"tgl_dibuat"`
	TglDiperbarui   time.Time `json:"tgl_diperbarui"`
	Status          string    `json:"status"`
	TicketCode      string    `json:"ticket_code"`
	Category        string    `json:"category"`
	Email           string    `json:"email"`
	AssignedTo      string    `json:"assigned_to"`
	Isi             string    `json:"isi" gorm:"->"`
	AreaCode        string    `json:"area_code"`
	AreaName        string    `json:"area_name" gorm:"->"`
	Regional        string    `json:"regional"`
	GrapariId       string    `json:"grapari_id"`
	GrapariName     string    `json:"grapari_name" gorm:"->"`
	TerminalId      string    `json:"terminal_id"`
	Lokasi          string    `json:"lokasi"`
	CategoryName    string    `json:"category_name" gorm:"->"`
	UserPembuat     string    `json:"user_pembuat" gorm:"->"`
	Assignee        string    `json:"assignee" gorm:"->"`
	SubCategory     string    `json:"sub_category"`
	StartTime       time.Time `json:"start_time" gorm:"->"`
	StartBy         string    `json:"start_by" gorm:"->"`
	CloseTime       time.Time `json:"close_time" gorm:"->"`
	CloseBy         string    `json:"close_by" gorm:"->"`
	AssigningTime   time.Time `json:"assigning_time"`
	AssigningBy     string    `json:"assigning_by"`
	AssignedToTeam  string    `json:"assigned_to_team"`
	TeamName        string    `json:"team_name" gorm:"->"`
	NoSPM           string    `json:"no_spm"`
	NoReqSPM        string    `json:"no_req_spm"`
	VisitStatus     string    `json:"visit_status"`
	CheckInTime     string    `json:"check_in_time" gorm:"->"`
	CheckOutTime    string    `json:"check_out_time" gorm:"->"`
}
