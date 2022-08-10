package model

import (
	"mime/multipart"
)

type HwReplacementRequest struct {
	TicketCode  string                `json:"ticket_code" form:"ticket_code"`
	HwId        string                `json:"hw_id" form:"hw_id"`
	OldSN       string                `json:"old_sn" form:"old_sn"`
	NewSN       string                `json:"new_sn" form:"new_sn"`
	Description string                `json:"description" form:"description"`
	Attachment  *multipart.FileHeader `json:"attachment" form:"attachment"`
	CreatedBy   string                `json:"created_by" form:"created_by"`
}
