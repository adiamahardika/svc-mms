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
	ResetPassword(request model.ResetPassword) (entity.User, error)
	Register(request model.RegisterRequest) (entity.User, error)
}

type userService struct {
	repository repository.UserRepositoryInterface
}

func UserService(repository repository.UserRepositoryInterface) *userService {
	return &userService{repository}
}

func (userService *userService) GetUser(request model.GetUserRequest) ([]entity.User, error) {
	user, error := userService.repository.GetUser(request)

	for index := range user {
		user[index].Password = "-"
	}
	
	return user, error
}

func (userService *userService) Login(request model.LoginRequest) ([]entity.User, error) {
	
	user, error := userService.repository.CheckUsername(request.Username)
	
	if (len(user) < 1) {
		error = fmt.Errorf("Username Not Found!")
		} else {
		error_check_pass := bcrypt.CompareHashAndPassword([]byte(user[0].Password), []byte(request.Password))
		if (error_check_pass != nil) {
			error = fmt.Errorf("Password Not Match")
		}
		user[0].Password = "-"
	}

	return user, error
}

func (userService *userService) ChangePassword(request model.ChangePassRequest) (entity.User, error) {
	var user entity.User
	date_now := time.Now()
	
	users, error := userService.repository.CheckUsername(request.Username)
	
	if (len(users) < 1) {
		error = fmt.Errorf("Username Not Found!")
		} else {
			
			error_check_pass := bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(request.OldPassword))
			if (error_check_pass != nil) {
				error = fmt.Errorf("Wrong Old Password!")
			} else {
				
				new_pass, error_hash_pass := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
				
				if (error_hash_pass != nil) {
					error = fmt.Errorf("There was an error creating new password!")
				} else {
					request.UpdatedAt = date_now
					request.NewPassword = string(new_pass)
	
					user, error = userService.repository.ChangePassword(request)
				}
			}
	}
	user.Password = "-"

	return user, error
}

func (userService *userService) ResetPassword(request model.ResetPassword) (entity.User, error) {
	var user entity.User
	date_now := time.Now()
	
	users, error := userService.repository.CheckUsername(request.Username)
	
	if (len(users) < 1) {
		error = fmt.Errorf("Username Not Found!")
		} else {
				
			new_pass, error_hash_pass := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
			
			if (error_hash_pass != nil) {
				error = fmt.Errorf("There was an error creating new password!")
			} else {
				new_request := model.ChangePassRequest{
					Username: request.Username,
					NewPassword: string(new_pass),
					UpdatedAt: date_now,
				}
	
				user, error = userService.repository.ChangePassword(new_request)
			}

			
	}
	user.Password = "-"

	return user, error
}

func (userService *userService) Register(request model.RegisterRequest) (entity.User, error) {
	var user entity.User
	date_now := time.Now()

	users, error := userService.repository.CheckUsername(request.Username)

	if (len(users) > 0) {
		error = fmt.Errorf("Username already exist!")
		} else {
			new_pass, error_hash_pass := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
			
			if (error_hash_pass != nil) {
				error = fmt.Errorf("There was an error creating new password!")
			} else {
				
			request.CreatedAt = date_now
			request.UpdatedAt = date_now
			request.Password = string(new_pass)
		
			user, error = userService.repository.Register(request)
		}
	}

	return user, error
}