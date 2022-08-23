package entity

import "time"

type Hardware struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	HwCode    string    `json:"hw_code"`
	IsActive  string    `json:"is_active"`
	CreatedBy string    `json:"created_by" form:"created_by"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}
