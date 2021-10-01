package service

import (
	"fmt"
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	GetUser(request model.GetUserRequest) ([]entity.User, error)
	Login(request model.LoginRequest) ([]entity.User, error)
}

type userService struct {
	repository repository.UserRepositoryInterface
}

func UserService(repository repository.UserRepositoryInterface) *userService {
	return &userService{repository}
}

func (userService *userService) GetUser(request model.GetUserRequest) ([]entity.User, error) {
	return userService.repository.GetUser(request)
}

func (userService *userService) Login(request model.LoginRequest) ([]entity.User, error) {
	
	user, error := userService.repository.CheckUsername(request)
	
	if (len(user) < 1) {
		error = fmt.Errorf("Username Not Found!")
		} else {
		check_pass := bcrypt.CompareHashAndPassword([]byte(user[0].Password), []byte(request.Password))
		if (check_pass != nil) {
			error = fmt.Errorf("Password Not Match")
		}
	}
	
	return user, error

}