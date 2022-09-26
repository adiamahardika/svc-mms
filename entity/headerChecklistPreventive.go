package entity

import "time"

type HeaderChecklistPreventive struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	PrevCode  string    `json:"prev_code"`
	Location  string    `json:"location"`
	UnitSn    string    `json:"unit_sn"`
	Date      string    `json:"date"`
	Time      string    `json:"time"`
	CreatedAt time.Time `json:"created_at"`
}
