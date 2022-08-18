package model

type GetAreaRequest struct {
	AreaCode []string `json:"area_code"`
	AreaName string   `json:"area_name"`
	Status   string   `json:"status"`
}
