package model

import "mime/multipart"

type UpdateTaskPreventiveRequest struct {
	PrevCode    string                `json:"prev_code" form:"prev_code"`
	Attachment  *multipart.FileHeader `json:"attachment" form:"attachment"`
	Description string                `json:"description" form:"description"`
	TaskName    string                `json:"task_name" form:"task_name"`
	Longitude   string                `json:"longitude" form:"longitude"`
	Latitude    string                `json:"latitude" form:"latitude"`
	AssignedBy  string                `json:"assigned_by" form:"assigned_by"`
	Status      string                `json:"status" form:"status"`
	Index       int                   `json:"index" form:"index"`
}
