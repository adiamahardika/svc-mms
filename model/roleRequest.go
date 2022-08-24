package model

import "svc-monitoring-maintenance/entity"

type CreateRoleRequest struct {
	Name          string                    `json:"name"`
	WebPermission []entity.MmsWebPermission `json:"web_permission" gorm:"foreignKey:Id"`
	AppPermission []entity.MmsAppPermission `json:"app_permission" gorm:"foreignKey:Id"`
}

type GetRoleRequest struct {
	IsActive string `json:"is_active"`
}
