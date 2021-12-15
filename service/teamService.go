package service

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
	"time"
)

type TeamServiceInterface interface {
	GetTeam() ([]entity.Team, error)
	CreateTeam(request model.CreateTeamRequest) (entity.Team, error)
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

func (teamService *teamService) CreateTeam(request model.CreateTeamRequest) (entity.Team, error) {
	date_now := time.Now()

	team_request := entity.Team{
		Name:      request.Name,
		CreatedAt: date_now,
	}

	_, error := teamService.repository.CreateTeam(team_request)

	return team_request, error
}
