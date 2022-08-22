package model

type GetReportResponse struct {
	Id              int    `json:"id" gorm:"primaryKey"`
	Judul           string `json:"judul"`
	UsernamePembuat string `json:"username_pembuat"`
	UpdatedBy       string `json:"updated_by" form:"updatedBy"`
	Prioritas       string `json:"prioritas"`
	TglDibuat       string `json:"tgl_dibuat"`
	TglDiperbarui   string `json:"tgl_diperbarui"`
	Status          string `json:"status"`
	TicketCode      string `json:"ticket_code"`
	Category        string `json:"category"`
	Email           string `json:"email"`
	AssignedTo      string `json:"assigned_to"`
	Isi             string `json:"isi" gorm:"->"`
	AreaCode        string `json:"area_code"`
	AreaName        string `json:"area_name" gorm:"->"`
	Regional        string `json:"regional"`
	GrapariId       string `json:"grapari_id"`
	GrapariName     string `json:"grapari_name" gorm:"->"`
	TerminalId      string `json:"terminal_id"`
	Lokasi          string `json:"lokasi"`
	CategoryName    string `json:"category_name" gorm:"->"`
	UserPembuat     string `json:"user_pembuat" gorm:"->"`
	Assignee        string `json:"assignee" gorm:"->"`
	SubCategory     string `json:"sub_category"`
	StartTime       string `json:"start_time" gorm:"->"`
	StartBy         string `json:"start_by" gorm:"->"`
	CloseTime       string `json:"close_time" gorm:"->"`
	CloseBy         string `json:"close_by" gorm:"->"`
	AssigningTime   string `json:"assigning_time"`
	AssigningBy     string `json:"assigning_by"`
	AssignedToTeam  string `json:"assigned_to_team"`
	TeamName        string `json:"team_name" gorm:"->"`
	NoSPM           string `json:"no_spm"`
	NoReqSPM        string `json:"no_req_spm"`
}
