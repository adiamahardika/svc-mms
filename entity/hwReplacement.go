package entity

import (
	"time"
)

type HwReplacement struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	TicketCode  string    `json:"ticket_code" form:"ticket_code"`
	HwId        string    `json:"hw_id" form:"hw_id"`
	HwName      string    `json:"hw_name" gorm:"->"`
	OldSN       string    `json:"old_sn" form:"old_sn"`
	NewSN       string    `json:"new_sn" form:"new_sn"`
	Description string    `json:"description" form:"description"`
	Attachment  string    `json:"attachment" form:"attachment"`
	CreatedBy   string    `json:"created_by" form:"created_by"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
}
