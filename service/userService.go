package service

import (
	"fmt"
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	GetUser(request model.GetUserRequest) ([]entity.User, error)
	Login(request model.LoginRequest) ([]entity.User, error)
	ChangePassword(request model.ChangePassRequest) (entity.User, error)
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
	
	user, error := userService.repository.CheckUsername(request.Username)
	
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

func (userService *userService) ChangePassword(request model.ChangePassRequest) (entity.User, error) {
	var user entity.User
	date_now := time.Now()
	
	request.UpdateAt = date_now
	
	users, error := userService.repository.CheckUsername(request.Username)
	
	if (len(users) < 1) {
		error = fmt.Errorf("Username Not Found!")
		} else {
			
			check_pass := bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(request.OldPassword))
			if (check_pass != nil) {
				error = fmt.Errorf("Wrong Old Password!")
			} else {
				user, error = userService.repository.ChangePassword(request)
			}
	}

	return user, error
}