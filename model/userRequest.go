package model

type GetUserRequest struct {
	Team string `json:"team"`
	Role string `json:"role"`
}