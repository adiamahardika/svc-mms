package entity

type Role struct {
	Id            int    `json:"id" gorm:"primaryKey"`
	Name          string `json:"name"`
	WebPermission string `json:"web_permission"`
	AppPermission string `json:"app_permission"`
}
