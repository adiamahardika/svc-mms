package service

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/repository"
)

type AppPermissionServiceInterface interface {
	GetAppPermission() ([]entity.MmsAppPermission, error)
}

type appPermissionService struct {
	appPermissionRepository repository.AppPermissionRepositoryInterface
}

func AppPermissionService(appPermissionRepository repository.AppPermissionRepositoryInterface) *appPermissionService {
	return &appPermissionService{appPermissionRepository}
}

func (appPermissionService *appPermissionService) GetAppPermission() ([]entity.MmsAppPermission, error) {

	app_permission, error := appPermissionService.appPermissionRepository.GetAppPermission()

	return app_permission, error
}
