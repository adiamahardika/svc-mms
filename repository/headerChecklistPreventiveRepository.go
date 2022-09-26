package repository

import "svc-monitoring-maintenance/entity"

type HeaderChecklistPreventiveRepositoryInterface interface {
	CreateHeaderChecklistPreventive(request *entity.HeaderChecklistPreventive) (entity.HeaderChecklistPreventive, error)
}

func (repo *repository) CreateHeaderChecklistPreventive(request *entity.HeaderChecklistPreventive) (entity.HeaderChecklistPreventive, error) {
	var headerChecklistPreventive entity.HeaderChecklistPreventive

	error := repo.db.Table("header_checklist_preventive").Create(&request).Find(&headerChecklistPreventive).Error

	return headerChecklistPreventive, error
}
