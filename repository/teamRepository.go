package repository

import "svc-monitoring-maintenance/entity"

type TeamRepositoryInterface interface {
	GetTeam() ([]entity.Team, error)
}

func (repo *repository) GetTeam() ([]entity.Team, error) {
	var team []entity.Team

	error := repo.db.Raw("SELECT * FROM team ORDER BY name").Find(&team).Error

	return team, error
}