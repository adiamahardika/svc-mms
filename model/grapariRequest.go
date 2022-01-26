package model

type GetGrapariRequest struct {
	Search   string `json:"search"`
	Area     string `json:"area"`
	Regional string `json:"regional"`
	Status   string `json:"status"`
}
