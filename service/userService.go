package service

import (
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
)

type UserServiceInterface interface {
	GetUser(request *model.GetUserRequest) ([]entity.User, error)
}

type userService struct {
	repository repository.UserRepositoryInterface
}

func UserService(repository repository.UserRepositoryInterface) *userService {
	return &userService{repository}
}

func (userService *userService) GetUser(request *model.GetUserRequest) ([]entity.User, error) {
	return userService.repository.GetUser(request)
}