package entity

import "time"

type TicketIsi struct {
	Id               int    `json:"id" gorm:"primaryKey"`
	UsernamePengirim string `json:"username_pengirim"`
	Isi              string `json:"isi"`
	KodeTicket       string `json:"kode_ticket"`
	Attachment1      string `json:"attachment1"`
	UrlAttachment1   string `json:"url_attachment1"`
	Attachment2      string `json:"attachment2"`
	UrlAttachment2   string `json:"url_attachment2"`
	TglDibuat        time.Time `json:"tgl_dibuat"`
}