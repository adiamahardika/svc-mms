package service

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/repository"
)

type WebPermissionServiceInterface interface {
	GetWebPermission() ([]entity.MmsWebPermission, error)
}

type webPermissionService struct {
	webPermissionRepository repository.WebPermissionRepositoryInterface
}

func WebPermissionService(webPermissionRepository repository.WebPermissionRepositoryInterface) *webPermissionService {
	return &webPermissionService{webPermissionRepository}
}

func (webPermissionService *webPermissionService) GetWebPermission() ([]entity.MmsWebPermission, error) {

	web_permission, error := webPermissionService.webPermissionRepository.GetWebPermission()

	return web_permission, error
}
