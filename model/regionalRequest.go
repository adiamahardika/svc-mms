package model

type GetRegionalRequest struct {
	Regional []string `json:"regional"`
	AreaCode []string `json:"area_code"`
	Status   string   `json:"status"`
}
