package service

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/repository"
	"time"
)

type MsChecklistHwServiceInterface interface {
	CreateMsChecklistHw(request *entity.MsChecklistHw) ([]entity.MsChecklistHw, error)
}

type msChecklistHwService struct {
	msChecklistHwRepository repository.MsChecklistHwRepositoryInterface
}

func MsChecklistHwService(msChecklistHwServiceInterface repository.MsChecklistHwRepositoryInterface) *msChecklistHwService {
	return &msChecklistHwService{msChecklistHwServiceInterface}
}

func (msChecklistHwService *msChecklistHwService) CreateMsChecklistHw(request *entity.MsChecklistHw) ([]entity.MsChecklistHw, error) {

	request.CreatedAt = time.Now()
	ms_checklist_hw, error := msChecklistHwService.msChecklistHwRepository.CreateMsChecklistHw(request)

	return ms_checklist_hw, error
}
