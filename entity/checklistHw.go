package entity

import "time"

type ChecklistHw struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Items     string    `json:"items"`
	Status    string    `json:"status"`
	PrevCode  string    `json:"prev_code"`
	CreatedAt time.Time `json:"created_at"`
}
