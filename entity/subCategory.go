package entity

import "time"

type SubCategory struct {
	Id         int       `json:"id" gorm:"primaryKey" gorm:"->"`
	Name       string    `json:"name"`
	IdCategory int       `json:"id_category"`
	Priority   string    `json:"priority"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
