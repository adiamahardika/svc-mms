package model

type GetGrapariRequest struct {
	Search    string   `json:"search"`
	Regional  []string `json:"regional"`
	AreaCode  []string `json:"areaCode"`
	GrapariId []string `json:"grapariId"`
	Status    string   `json:"status"`
}
