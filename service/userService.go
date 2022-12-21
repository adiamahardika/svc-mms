package service

import (
	"fmt"
	"strconv"
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	GetUser(request model.GetUserRequest) ([]model.GetUserResponse, error)
	ChangePassword(request model.ChangePassRequest) (model.GetUserResponse, error)
	ResetPassword(request model.ResetPassword) (model.GetUserResponse, error)
	GetDetailUser(request string) ([]model.GetUserResponse, error)
	UpdateKeyHp(request *model.UpdateKeyHpRequest) (model.UpdateKeyHpRequest, error)
	UpdateUser(request *entity.User) (entity.User, error)
}

type userService struct {
	repository repository.UserRepositoryInterface
}

func UserService(repository repository.UserRepositoryInterface) *userService {
	return &userService{repository}
}

func (userService *userService) GetUser(request model.GetUserRequest) ([]model.GetUserResponse, error) {
	user, error := userService.repository.GetUser(request)

	return user, error
}

func (userService *userService) ChangePassword(request model.ChangePassRequest) (model.GetUserResponse, error) {
	var user model.GetUserResponse
	date_now := time.Now()

	users, error := userService.repository.CheckUsername(request.Username)

	if len(users) < 1 {
		error = fmt.Errorf("Username Not Found!")
	} else {

		error_check_pass := bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(request.OldPassword))
		if error_check_pass != nil {
			error = fmt.Errorf("Wrong Old Password!")
		} else {

			new_pass, error_hash_pass := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)

			if error_hash_pass != nil {
				error = fmt.Errorf("There was an error creating new password!")
			} else {
				request.UpdatedAt = date_now
				request.NewPassword = string(new_pass)

				user, error = userService.repository.ChangePassword(request)
			}
		}
	}

	return user, error
}

func (userService *userService) ResetPassword(request model.ResetPassword) (model.GetUserResponse, error) {
	var user model.GetUserResponse
	date_now := time.Now()

	users, error := userService.repository.CheckUsername(request.Username)

	if len(users) < 1 {
		error = fmt.Errorf("Username Not Found!")
	} else {

		new_pass, error_hash_pass := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)

		if error_hash_pass != nil {
			error = fmt.Errorf("There was an error creating new password!")
		} else {
			new_request := model.ChangePassRequest{
				Username:    request.Username,
				NewPassword: string(new_pass),
				UpdatedAt:   date_now,
			}

			user, error = userService.repository.ChangePassword(new_request)
		}

	}

	return user, error
}

func (userService *userService) GetDetailUser(request string) ([]model.GetUserResponse, error) {

	user_id, _ := strconv.Atoi(request)
	user, error := userService.repository.GetDetailUser(user_id)

	return user, error
}

func (userService *userService) UpdateKeyHp(request *model.UpdateKeyHpRequest) (model.UpdateKeyHpRequest, error) {

	key_hp, error := userService.repository.UpdateKeyHp(&model.LoginRequest{
		Username: request.Username,
		KeyHp:    request.KeyHp,
	})

	request.KeyHp = key_hp

	return *request, error
}

func (userService *userService) UpdateUser(request *entity.User) (entity.User, error) {

	request.UpdatedAt = time.Now()

	user, error := userService.repository.UpdateUser(request)

	return user, error
}
