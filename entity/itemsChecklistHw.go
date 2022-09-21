package entity

import "time"

type ItemsChecklistHw struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}
