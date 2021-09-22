package entity

import "time"

type TaskList struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	KodeTicket  string `json:"kode_ticket"`
	Attachment  string `json:"attachment"`
	Description string `json:"description"`
	TaskName    string `json:"task_name"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	AssignedBy  string `json:"assigned_by"`
	Status      string `json:"status"`
	CreatedBy   time.Time `json:"created_by"`
}