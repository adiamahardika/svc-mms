package entity

import "time"

type User struct {
	Id        int    	`json:"id" gorm:"primaryKey"`
	Name      string 	`json:"name"`
	Username  string 	`json:"username"`
	Password  string 	`json:"password"`
	Email     string 	`json:"email"`
	Gender    string 	`json:"gender"`
	Team 	  string 	`json:"team"`
	TeamName  string 	`json:"team_name"`
	Role 	  string 	`json:"role"`
	RoleName  string 	`json:"role_name"`
	UpdatedAt time.Time	`json:"updated_at"`
	CreatedAt time.Time	`json:"created_at"`
}