package entity

import "time"

type Preventive struct {
	Id             int       `json:"id" gorm:"primaryKey"`
	PrevCode       string    `json:"prev_code"`
	VisitDate      string    `json:"visit_date"`
	AreaCode       string    `json:"area_code"`
	AreaName       string    `json:"area_name" gorm:"->"`
	Regional       string    `json:"regional"`
	GrapariId      string    `json:"grapari_id"`
	GrapariName    string    `json:"grapari_name" gorm:"->"`
	Location       string    `json:"location"`
	TerminalId     string    `json:"terminal_id"`
	AssignedTo     string    `json:"assigned_to"`
	UserName       string    `json:"user_name" gorm:"->"`
	AssignedToTeam string    `json:"assigned_to_team"`
	TeamName       string    `json:"team_name" gorm:"->"`
	Status         string    `json:"status"`
	CreatedBy      string    `json:"created_by"`
	Creator        string    `json:"creator" gorm:"->"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedBy      string    `json:"updated_by"`
	UpdatedAt      time.Time `json:"updated_at"`
	NoSPM          string    `json:"no_spm"`
	NoReqSPM       string    `json:"no_req_spm"`
	Note           string    `json:"note"`
	Judul          string    `json:"judul"`
	Email          string    `json:"email"`
	CheckInTime    string    `json:"check_in_time" gorm:"->"`
	CheckOutTime   string    `json:"check_out_time" gorm:"->"`
}
