package repository

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
)

type UserRepositoryInterface interface {
	GetUser(request model.GetUserRequest) ([]model.GetUserResponse, error)
	CheckUsername(request string) ([]entity.User, error)
	ChangePassword(request model.ChangePassRequest) (model.GetUserResponse, error)
	GetDetailUser(request int) ([]model.GetUserResponse, error)
	UpdateKeyHp(request *model.LoginRequest) (string, error)
	UpdateUser(request *entity.User) (entity.User, error)
}

func (repo *repository) GetUser(request model.GetUserRequest) ([]model.GetUserResponse, error) {
	var user []model.GetUserResponse

	error := repo.db.Raw("SELECT users.*, role.name as role_name, team.name as team_name FROM users LEFT OUTER JOIN role ON (users.role = CAST(role.id AS varchar(10))) LEFT OUTER JOIN team ON (users.team = CAST(team.id AS varchar(10))) WHERE users.role LIKE @Role AND users.team LIKE @Team AND users.employment_status ILIKE @EmploymentStatus AND users.is_active = 'true' ORDER BY users.name", model.GetUserRequest{
		Team:             "%" + request.Team + "%",
		Role:             "%" + request.Role + "%",
		EmploymentStatus: "%" + request.EmploymentStatus + "%",
	}).Find(&user).Error

	return user, error
}

func (repo *repository) CheckUsername(request string) ([]entity.User, error) {
	var user []entity.User

	error := repo.db.Raw("SELECT users.*, role.name as role_name, team.name as team_name FROM users LEFT OUTER JOIN role ON (users.role = CAST(role.id AS varchar(10))) LEFT OUTER JOIN team ON (users.team = CAST(team.id AS varchar(10))) WHERE username = @Username", model.LoginRequest{
		Username: request,
	}).Find(&user).Error

	return user, error
}

func (repo *repository) ChangePassword(request model.ChangePassRequest) (model.GetUserResponse, error) {
	var user model.GetUserResponse

	error := repo.db.Raw("UPDATE users SET password = @NewPassword, updated_at = @UpdatedAt WHERE username = @Username RETURNING users.*", model.ChangePassRequest{
		Username:    request.Username,
		NewPassword: request.NewPassword,
		UpdatedAt:   request.UpdatedAt,
	}).Find(&user).Error

	return user, error
}

func (repo *repository) GetDetailUser(request int) ([]model.GetUserResponse, error) {
	var user []model.GetUserResponse

	error := repo.db.Raw("SELECT users.*, role.name as role_name, team.name as team_name FROM users LEFT OUTER JOIN role ON (users.role = CAST(role.id AS varchar(10))) LEFT OUTER JOIN team ON (users.team = CAST(team.id AS varchar(10))) WHERE users.id = @Id", entity.User{
		Id: request,
	}).Find(&user).Error

	return user, error
}

func (repo *repository) UpdateKeyHp(request *model.LoginRequest) (string, error) {
	var user entity.User

	error := repo.db.Raw("UPDATE users SET key_hp = @KeyHp WHERE username = @Username RETURNING users.*", model.LoginRequest{
		KeyHp:    request.KeyHp,
		Username: request.Username,
	}).Find(&user).Error

	return user.KeyHp, error
}

func (repo *repository) UpdateUser(request *entity.User) (entity.User, error) {
	var user entity.User

	error := repo.db.Raw("UPDATE users SET email = @Email, role = @Role, team = @Team, name = @Name, updated_at = @UpdatedAt, is_active = @IsActive WHERE id = @Id RETURNING users.*", request).Find(&user).Error

	return user, error
}
