package entity

type Role struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}