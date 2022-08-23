package model

import "svc-monitoring-maintenance/entity"

type GetRoleResponse struct {
	Id            int                        `json:"id" gorm:"primaryKey"`
	Name          string                     `json:"name"`
	WebPermission []*entity.MmsWebPermission `json:"web_permission" gorm:"foreignKey:Id"`
	AppPermission []*entity.MmsAppPermission `json:"app_permission" gorm:"foreignKey:Id"`
}
