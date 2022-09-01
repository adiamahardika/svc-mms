package entity

import (
	"time"
)

type HwReplacement struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	TicketCode  string    `json:"ticket_code"`
	HwId        string    `json:"hw_id"`
	HwName      string    `json:"hw_name" gorm:"->"`
	OldSN       string    `json:"old_sn"`
	NewSN       string    `json:"new_sn"`
	Description string    `json:"description"`
	Attachment  string    `json:"attachment"`
	StatusId    string    `json:"status_id"`
	Status      string    `json:"status" gorm:"->"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}
