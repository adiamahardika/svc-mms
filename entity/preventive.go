package entity

import "time"

type Preventive struct {
	Id         int       `json:"id" gorm:"primaryKey"`
	PrevCode   string    `json:"prev_code"`
	VisitDate  string    `json:"visit_date"`
	TerminalId string    `json:"terminal_id"`
	AssignedTo string    `json:"assigned_to"`
	Status     string    `json:"status"`
	CreatedBy  string    `json:"created_by"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedBy  string    `json:"updated_by"`
	UpdatedAt  time.Time `json:"updated_at"`
}