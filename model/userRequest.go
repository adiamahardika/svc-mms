package model

import "time"

type GetUserRequest struct {
	Team string `json:"team"`
	Role string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangePassRequest struct {
	Username    string 		`json:"username"`
	OldPassword string 		`json:"old_password"`
	NewPassword string 		`json:"new_password"`
	UpdatedAt    time.Time  `json:"updated_at"`
}