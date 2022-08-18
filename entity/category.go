package entity

import "time"

type Category struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	SubCategory string    `json:"sub_category" gorm:"->"`
	IsActive    string    `json:"is_active"`
	UpdateAt    time.Time `json:"update_at"`
}
