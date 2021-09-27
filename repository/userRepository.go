package repository

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type UserRepositoryInterface interface {
	GetUser(request *model.GetUserRequest) ([]entity.User, error)
}

func (repo *repository) GetUser(request *model.GetUserRequest) ([]entity.User, error){
	var user []entity.User

	error := repo.db.Raw("SELECT users.*, role.name as role_name, team.name as team_name FROM users LEFT OUTER JOIN role ON (users.role = CAST(role.id AS varchar(10))) LEFT OUTER JOIN team ON (users.team = CAST(team.id AS varchar(10))) WHERE users.role LIKE @Role AND users.team LIKE @Team ORDER BY users.name", model.GetUserRequest{
		Team: "%" + request.Team + "%",
		Role: "%" + request.Role + "%",
	}).Find(&user).Error

	return user, error
}