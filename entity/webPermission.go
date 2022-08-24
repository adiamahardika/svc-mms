package entity

type MmsWebPermission struct {
	Id             int    `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	PermissionCode string `json:"permission_code"`
}
