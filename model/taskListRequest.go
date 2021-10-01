package model

import (
	"mime/multipart"
)

type GetTaskListRequest struct {
	TicketCode string `json:"ticket_code"`
}

type UpdateTaskListRequest struct {
	TicketCode  string    `json:"ticket_code" form:"ticket_code"`
	Attachment  *multipart.FileHeader `json:"attachment" form:"attachment"`
	Description string    `json:"description" form:"description"`
	TaskName    string    `json:"task_name" form:"task_name"`
	Longitude   string    `json:"longitude" form:"longitude"`
	Latitude    string    `json:"latitude" form:"latitude"`
	AssignedBy  string    `json:"assigned_by" form:"assigned_by"`
	Status      string    `json:"status" form:"status"`
}