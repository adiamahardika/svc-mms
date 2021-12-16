package model

type CreateTeamRequest struct {
	Name string `json:"name"`
}

type GetTeamRequest struct {
	IsActive string `json:"is_active"`
}
