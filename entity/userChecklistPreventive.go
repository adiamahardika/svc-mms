package entity

import "time"

type UserChecklistPreventive struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	PrevCode  string    `json:"prev_code"`
	UserType  string    `json:"user_type"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	Date      string    `json:"date"`
	Time      string    `json:"time"`
	Signature string    `json:"signature"`
	CreatedAt time.Time `json:"created_at"`
}
