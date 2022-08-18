package service

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
)

type GrapariServiceInterface interface {
	GetGrapari(request model.GetGrapariRequest) ([]entity.MsGrapari, error)
}

type grapariService struct {
	repository repository.GrapariRepositoryInterface
}

func GrapariService(repository repository.GrapariRepositoryInterface) *grapariService {
	return &grapariService{repository}
}

func (grapariService *grapariService) GetGrapari(request model.GetGrapariRequest) ([]entity.MsGrapari, error) {
	terminal, error := grapariService.repository.GetGrapari(request)

	return terminal, error
}
