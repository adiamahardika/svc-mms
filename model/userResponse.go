package model

import "time"

type GetUserResponse struct {
	Id               int                `json:"id" gorm:"primaryKey"`
	Name             string             `json:"name"`
	Username         string             `json:"username"`
	Email            string             `json:"email"`
	Team             string             `json:"team"`
	TeamName         string             `json:"team_name"`
	Role             []*GetRoleResponse `json:"role" gorm:"foreignKey:Id"`
	RoleName         string             `json:"role_name"`
	KeyHp            string             `json:"key_hp"`
	Nik              string             `json:"nik"`
	Position         string             `json:"position"`
	Department       string             `json:"department"`
	IsActive         string             `json:"is_active"`
	EmploymentStatus string             `json:"employment_status"`
	UpdatedAt        time.Time          `json:"updated_at"`
	CreatedAt        time.Time          `json:"created_at"`
}
