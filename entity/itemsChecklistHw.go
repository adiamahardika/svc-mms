package entity

import "time"

type ItemsChecklistHw struct {
	Id            int       `json:"id" gorm:"primaryKey"`
	Name          string    `json:"items"`
	DefaultStatus string    `json:"status"`
	CreatedBy     string    `json:"created_by"`
	CreatedAt     time.Time `json:"created_at"`
}
