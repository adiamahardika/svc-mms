package entity

type RoleHasAppPermission struct {
	Id           int `json:"id" gorm:"primaryKey"`
	IdRole       int `json:"id_role"`
	IdPermission int `json:"id_permission"`
}
