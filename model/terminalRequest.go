package model

type GetTerminalRequest struct {
	Search        string `json:"search"`
	GrapariId     string `json:"grapari_id"`
	Area          string `json:"area"`
	Regional      string `json:"regional"`
	CtpType       string `json:"ctp_type"`
	Kecamatan     string `json:"kecamatan"`
	KotaKabupaten string `json:"kota_kabupaten"`
}
