package model

import (
	"mime/multipart"
)

type CreateHwReplacementRequest struct {
	TicketCode    string                `json:"ticket_code" form:"ticket_code"`
	HwId          string                `json:"hw_id" form:"hw_id"`
	OldSN         string                `json:"old_sn" form:"old_sn"`
	NewSN         string                `json:"new_sn" form:"new_sn"`
	Description   string                `json:"description" form:"description"`
	StatusId      string                `json:"status_id" form:"status_id"`
	OldAttachment *multipart.FileHeader `json:"old_attachment" form:"old_attachment"`
	NewAttachment *multipart.FileHeader `json:"new_attachment" form:"new_attachment"`
	CreatedBy     string                `json:"created_by" form:"created_by"`
}

type GetHwReplacementRequest struct {
	TicketCode string `json:"ticket_code" form:"ticket_code"`
}
