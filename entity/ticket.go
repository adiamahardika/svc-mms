package entity

import "time"

type Ticket struct {
	Id               int 		`json:"id" gorm:"primaryKey"`
	Judul            string 	`json:"judul"`
	UsernamePembuat  string		`json:"username_pembuat"`
	UsernamePembalas string 	`json:"username_pembalas"`
	Prioritas        string		`json:"prioritas"`
	TglDibuat        time.Time	`json:"tgl_dibuat"`
	TglDiperbarui    time.Time	`json:"tgl_diperbarui"`
	TotalWaktu		 string		`json:"total_waktu"`
	Status			 string		`json:"status"`
	KodeTicket		 string		`json:"kode_ticket"`
	Kategori		 string		`json:"kategori"`
	Lokasi			 string		`json:"lokasi"`
	TerminalId		 string		`json:"terminal_id"`
	Email			 string		`json:"email"`
	AssignedTo		 string		`json:"assigned_to"`
	AssignedToTeam	 string		`json:"assigned_to_team"`
}