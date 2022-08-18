package entity

import "time"

type MsRegional struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Regional  string    `json:"regional"`
	Area      string    `json:"area"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
