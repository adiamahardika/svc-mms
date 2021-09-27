package entity

import "time"

type Team struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}