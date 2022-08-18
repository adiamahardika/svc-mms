package model

import (
	"svc-monitoring-maintenance/entity"
	"time"
)

type CreateCategoryRequest struct {
	Id          int                  `json:"id" gorm:"primaryKey"`
	Name        string               `json:"name"`
	SubCategory []entity.SubCategory `json:"sub_category" gorm:"foreignKey:Id"`
	IsActive    string               `json:"is_active"`
	UpdateAt    time.Time            `json:"update_at"`
}

type GetCategoryRequest struct {
	Size       int    `json:"size"`
	PageNo     int    `json:"page_no"`
	StartIndex int    `json:"start_index"`
	SortBy     string `json:"sort_by"`
	OrderBy    string `json:"order_by"`
	IsActive   string `json:"is_active"`
}
