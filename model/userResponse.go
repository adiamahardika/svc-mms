package model

import "time"

type GetUserResponse struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Team      string    `json:"team"`
	TeamName  string    `json:"team_name"`
	Role      string    `json:"role"`
	RoleName  string    `json:"role_name"`
	KeyHp     string    `json:"key_hp"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
