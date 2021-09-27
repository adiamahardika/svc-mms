package service

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/repository"
)

type TeamServiceInterface interface {
	GetTeam() ([]entity.Team, error)
}

type teamService struct {
	repository repository.TeamRepositoryInterface
}

func TeamService(repository repository.TeamRepositoryInterface) *teamService {
	return &teamService{repository}
}

func (teamService *teamService) GetTeam() ([]entity.Team, error) {
	return teamService.repository.GetTeam()
}