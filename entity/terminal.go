package entity

import "time"

type Terminal struct {
	Id               int       `json:"id" gorm:"primaryKey"`
	TerminalId       string    `json:"terminal_id"`
	TerminalName     string    `json:"terminal_name"`
	Area             string    `json:"area"`
	Regional         string    `json:"regional"`
	CtpType          string    `json:"ctp_type"`
	TerminalLocation string    `json:"terminal_location"`
	Kecamatan        string    `json:"kecamatan"`
	KotaKabupaten    string    `json:"kota_kabupaten"`
	KodePos          string    `json:"kode_pos"`
	Pic              string    `json:"pic"`
	KontakPic        string    `json:"kontak_pic"`
	CreatedBy        string    `json:"created_by"`
	CreatedDate      time.Time `json:"created_date"`
	UpdatedBy        string    `json:"created_by"`
	UpdatedDate      string    `json:"updated_date"`
	Status           string    `json:"status"`
	RecStatus        string    `json:"rec_status"`
	RegionalCode     string    `json:"regional_code"`
	GrapariId        string    `json:"grapari_id"`
	Latitude         string    `json:"latitude"`
	Longitude        string    `json:"longitude"`
}
