package entity

import "time"

type MsChecklistHw struct {
	Name      string    `json:"name"`
	HWCode    string    `json:"hw_code"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}
