package service

import (
	"fmt"
	"os"
	"strconv"
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	GetUser(request model.GetUserRequest) ([]model.GetUserResponse, error)
	Login(request model.LoginRequest) (model.GetUserResponse, model.LoginResponse, error)
	ChangePassword(request model.ChangePassRequest) (model.GetUserResponse, error)
	ResetPassword(request model.ResetPassword) (model.GetUserResponse, error)
	Register(request model.RegisterRequest) (entity.User, error)
	GetDetailUser(request string) ([]model.GetUserResponse, error)
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

func (userService *userService) Login(request model.LoginRequest) (model.GetUserResponse, model.LoginResponse, error) {
	var user_response model.GetUserResponse
	var login_response model.LoginResponse
	user, error := userService.repository.CheckUsername(request.Username)

	if len(user) < 1 {
		error = fmt.Errorf("Username Not Found!")
	} else {
		error_check_pass := bcrypt.CompareHashAndPassword([]byte(user[0].Password), []byte(request.Password))

		if error_check_pass != nil {
			error = fmt.Errorf("Password Not Match")
		}
		user_response = model.GetUserResponse{
			Id:        user[0].Id,
			Name:      user[0].Name,
			Username:  user[0].Username,
			Email:     user[0].Email,
			Team:      user[0].Team,
			TeamName:  user[0].TeamName,
			Role:      user[0].Role,
			RoleName:  user[0].RoleName,
			UpdatedAt: user[0].UpdatedAt,
			CreatedAt: user[0].CreatedAt,
		}

		expirationTime := time.Now().Add(time.Minute * 5)
		claims := &model.Claims{
			Username: request.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		jwtKey := []byte(os.Getenv("API_SECRET"))
		tokenString, err := token.SignedString(jwtKey)

		if err != nil {
			error = err
		}

		login_response = model.LoginResponse{
			Token: tokenString,
		}
	}

	return user_response, login_response, error
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

func (userService *userService) Register(request model.RegisterRequest) (entity.User, error) {
	var user entity.User
	date_now := time.Now()

	users, error := userService.repository.CheckUsername(request.Username)

	if len(users) > 0 {
		error = fmt.Errorf("Username already exist!")
	} else {
		new_pass, error_hash_pass := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

		if error_hash_pass != nil {
			error = fmt.Errorf("There was an error creating new password!")
		} else {

			request.CreatedAt = date_now
			request.UpdatedAt = date_now
			request.Password = string(new_pass)
			request.Changepass = "0"

			_, error = userService.repository.Register(request)
			users, error = userService.repository.CheckUsername(request.Username)
			user = users[0]
		}
	}

	return user, error
}

func (userService *userService) GetDetailUser(request string) ([]model.GetUserResponse, error) {

	user_id, _ := strconv.Atoi(request)
	user, error := userService.repository.GetDetailUser(user_id)

	return user, error
}
