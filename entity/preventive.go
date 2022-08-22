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
	UserName       string    `json:"user_name"`
	AssignedToTeam string    `json:"assigned_to_team"`
	TeamName       string    `json:"team_name" gorm:"->"`
	Status         string    `json:"status"`
	CreatedBy      string    `json:"created_by"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedBy      string    `json:"updated_by"`
	UpdatedAt      time.Time `json:"updated_at"`
}
