package model

type GetPreventiveRequest struct {
	Search         string   `json:"search"`
	Status         []string `json:"status"`
	AssignedTo     string   `json:"assigned_to"`
	AssignedToTeam string   `json:"assigned_to_team"`
	StartDate      string   `json:"start_date" binding:"required"`
	EndDate        string   `json:"end_date" binding:"required"`
	AreaCode       []string `json:"area_code"`
	Regional       []string `json:"regional"`
	GrapariId      []string `json:"grapari_id"`
	PageNo         int      `json:"page_no"`
	PageSize       int      `json:"page_size"`
	StartIndex     int      `json:"start_index"`
}

type CountPreventiveByStatusRequest struct {
	AssignedTo     string `json:"assigned_to"`
	AssignedToTeam string `json:"assigned_to_team"`
	StartDate      string `json:"start_date" binding:"required"`
	EndDate        string `json:"end_date" binding:"required"`
}
