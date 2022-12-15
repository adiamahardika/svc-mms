package model

import "time"

type GetUserRequest struct {
	Team             string `json:"team"`
	Role             string `json:"role"`
	EmploymentStatus string `json:"employment_status"`
}

type ChangePassRequest struct {
	Username    string    `json:"username"`
	OldPassword string    `json:"old_password"`
	NewPassword string    `json:"new_password"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ResetPassword struct {
	Username    string    `json:"username"`
	NewPassword string    `json:"new_password"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RegisterRequest struct {
	Name             string    `json:"name"`
	Username         string    `json:"username"`
	Password         string    `json:"password"`
	Changepass       string    `json:"changepass"`
	Email            string    `json:"email"`
	Team             string    `json:"team"`
	Role             string    `json:"role"`
	Nik              string    `json:"nik"`
	Position         string    `json:"position"`
	Department       string    `json:"department"`
	IsActive         string    `json:"is_active"`
	EmploymentStatus string    `json:"employment_status"`
	UpdatedAt        time.Time `json:"updated_at"`
	CreatedAt        time.Time `json:"created_at"`
}

type UpdateKeyHpRequest struct {
	Username string `json:"username" binding:"required"`
	KeyHp    string `json:"key_hp" binding:"required"`
}
