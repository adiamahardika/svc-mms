package repository

import "svc-monitoring-maintenance/entity"

type TeamRepositoryInterface interface {
	GetTeam() ([]entity.Team, error)
	CreateTeam(request entity.Team) (entity.Team, error)
	UpdateTeam(request entity.Team) (entity.Team, error)
}

func (repo *repository) GetTeam() ([]entity.Team, error) {
	var team []entity.Team

	error := repo.db.Raw("SELECT * FROM team ORDER BY name").Find(&team).Error

	return team, error
}

func (repo *repository) CreateTeam(request entity.Team) (entity.Team, error) {
	var team entity.Team

	error := repo.db.Table("team").Create(&request).Error

	return team, error
}

func (repo *repository) UpdateTeam(request entity.Team) (entity.Team, error) {
	var team entity.Team

	error := repo.db.Raw("UPDATE team SET name = @Name WHERE id = @Id RETURNING team.*", request).Find(&team).Error

	return team, error
}
