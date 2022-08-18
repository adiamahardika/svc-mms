package entity

import "time"

type MsGrapari struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	GrapariId string    `json:"grapari_id"`
	Name      string    `json:"name"`
	Area      string    `json:"area"`
	Regional  string    `json:"regional"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    string    `json:"status"`
}
