package repository

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type TeamRepositoryInterface interface {
	GetTeam(request model.GetTeamRequest) ([]entity.Team, error)
	CreateTeam(request entity.Team) (entity.Team, error)
	UpdateTeam(request entity.Team) (entity.Team, error)
	DeleteTeam(Id int) error
}

func (repo *repository) GetTeam(request model.GetTeamRequest) ([]entity.Team, error) {
	var team []entity.Team

	error := repo.db.Raw("SELECT * FROM team WHERE is_active LIKE @IsActive ORDER BY name", model.GetTeamRequest{
		IsActive: "%" + request.IsActive + "%",
	}).Find(&team).Error

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

func (repo *repository) DeleteTeam(Id int) error {
	var team entity.Team

	error := repo.db.Raw("UPDATE team SET is_active = ? WHERE id = ? RETURNING team.*", "false", Id).Find(&team).Error

	return error
}
