package entity

type MmsAppPermission struct {
	Id             int    `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	PermissionCode string `json:"permission_code"`
}
