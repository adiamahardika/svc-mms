package service

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
)

type TerminalServiceInterface interface {
	GetTerminal(request model.GetTerminalRequest) ([]entity.MsTerminal, error)
}

type terminalService struct {
	repository repository.TerminalRepositoryInterface
}

func TerminalService(repository repository.TerminalRepositoryInterface) *terminalService {
	return &terminalService{repository}
}

func (terminalService *terminalService) GetTerminal(request model.GetTerminalRequest) ([]entity.MsTerminal, error) {
	list_terminal, error := terminalService.repository.GetTerminal(request)

	return list_terminal, error
}
