package model

import "time"

type GetUserRequest struct {
	Team string `json:"team"`
	Role string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ChangePassRequest struct {
	Username    string 		`json:"username"`
	OldPassword string 		`json:"old_password"`
	NewPassword string 		`json:"new_password"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type ResetPassword struct {
	Username    string 		`json:"username"`
	NewPassword string 		`json:"new_password"`
	UpdatedAt   time.Time 	`json:"updated_at"`
}

type RegisterRequest struct {
	Name      string 	`json:"name"`
	Username  string 	`json:"username"`
	Password  string 	`json:"password"`
	Changepass  string 	`json:"changepass"`
	Email     string 	`json:"email"`
	Team 	  string 	`json:"team"`
	Role 	  string 	`json:"role"`
	UpdatedAt time.Time	`json:"updated_at"`
	CreatedAt time.Time	`json:"created_at"`
}