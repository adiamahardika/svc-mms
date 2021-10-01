package entity

import "time"

type TaskList struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	TicketCode  string `json:"ticket_code"`
	Attachment  string `json:"attachment"`
	Description string `json:"description"`
	TaskName    string `json:"task_name"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	AssignedBy  string `json:"assigned_by"`
	Status      string `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}