package model

type GetUserRequest struct {
	Team string `json:"team"`
	Role string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}