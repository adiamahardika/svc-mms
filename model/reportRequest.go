package model

type GetReportRequest struct {
	AssignedTo      string   `json:"assigned_to"`
	AssignedToTeam  string   `json:"assigned_to_team"`
	Category        []string `json:"category"`
	AreaCode        []string `json:"area_code"`
	Regional        []string `json:"regional"`
	GrapariId       []string `json:"grapari_id"`
	Priority        []string `json:"priority"`
	Status          []string `json:"status"`
	UsernamePembuat []string `json:"username_pembuat"`
	StartDate       string   `json:"start_date" binding:"required"`
	EndDate         string   `json:"end_date" binding:"required"`
}
