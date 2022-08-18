package entity

import "time"

type MsArea struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	AreaCode  string    `json:"area_code"`
	AreaName  string    `json:"area_name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
