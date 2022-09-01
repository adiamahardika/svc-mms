package entity

type HwReplacementStatus struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Status   string `json:"status"`
	IsActive string `json:"is_active"`
}
