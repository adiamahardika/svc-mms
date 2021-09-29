package model

import (
	"mime/multipart"
)

type GetTaskListRequest struct {
	KodeTicket string `json:"kode_ticket"`
}

type UpdateTaskListRequest struct {
	KodeTicket  string    `json:"kode_ticket" form:"kode_ticket"`
	Attachment  *multipart.FileHeader `json:"attachment" form:"attachment"`
	Description string    `json:"description" form:"description"`
	TaskName    string    `json:"task_name" form:"task_name"`
	Longitude   string    `json:"longitude" form:"longitude"`
	Latitude    string    `json:"latitude" form:"latitude"`
	AssignedBy  string    `json:"assigned_by" form:"assigned_by"`
	Status      string    `json:"status" form:"status"`
}