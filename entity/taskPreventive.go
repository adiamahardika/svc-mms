package entity

import "time"

type TaskPreventive struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	PrevCode    string    `json:"prev_code"`
	Attachment  string    `json:"attachment"`
	Description string    `json:"description"`
	TaskName    string    `json:"task_name"`
	Longitude   string    `json:"longitude"`
	Latitude    string    `json:"latitude"`
	AssignedBy  string    `json:"assigned_by"`
	Status      string    `json:"status"`
	Index       int       `json:"index"`
	CreatedAt   time.Time `json:"created_at"`
}
