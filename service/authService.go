package service

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/general"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	Login(request model.LoginRequest) (model.GetUserResponse, model.LoginResponse, error)
	Register(request model.RegisterRequest) (entity.User, error)
}

type authService struct {
	userRepository repository.UserRepositoryInterface
	authRepository repository.AuthRepositoryInterface
}

func AuthService(userRepository repository.UserRepositoryInterface, authRepository repository.AuthRepositoryInterface) *authService {
	return &authService{userRepository, authRepository}
}

func Authentication() gin.HandlerFunc {
	return func(context *gin.Context) {

		token_string := context.Request.Header.Get("token")
		claims := &model.Claims{}
		description := []string{}
		jwtKey := []byte(os.Getenv("API_SECRET"))
		var status model.StandardResponse

		token, error := jwt.ParseWithClaims(token_string, claims,
			func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

		if token == nil {
			error = fmt.Errorf(fmt.Sprintf("Please provide token!"))
		} else if time.Now().Unix() > claims.ExpiresAt {
			error = fmt.Errorf(fmt.Sprintf("Your token is expired!"))
		} else if error != nil {
			error = fmt.Errorf(fmt.Sprintf("Your token is invalid!"))
		}

		if error != nil {
			description = append(description, error.Error())
			status = model.StandardResponse{
				HttpStatus:  http.StatusUnauthorized,
				StatusCode:  general.ErrorStatusCode,
				Description: description,
			}
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": status,
			})
			context.Abort()
		}
		context.Next()

	}
}

func (authService *authService) Authorization() gin.HandlerFunc {
	return func(context *gin.Context) {

		signature_key := context.Request.Header.Get("signature-key")
		token_string := context.Request.Header.Get("token")
		claims := &model.Claims{}
		description := []string{}
		jwtKey := []byte(os.Getenv("API_SECRET"))
		var status model.StandardResponse

		_, error := jwt.ParseWithClaims(token_string, claims,
			func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

		user, error := authService.userRepository.CheckUsername(claims.Username)
		generate_sk := general.GetMD5Hash(claims.Username, strconv.Itoa(user[0].Id))

		if signature_key == "" {
			error = fmt.Errorf(fmt.Sprintf("Please provide signature-key!"))
		} else if signature_key != generate_sk {
			error = fmt.Errorf(fmt.Sprintf("Your signature-key is invalid!"))
		}

		if error != nil {
			description = append(description, error.Error())
			status = model.StandardResponse{
				HttpStatus:  http.StatusUnauthorized,
				StatusCode:  general.ErrorStatusCode,
				Description: description,
			}
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": status,
			})
			context.Abort()
		}
		context.Next()

	}
}

func (authService *authService) Login(request model.LoginRequest) (model.GetUserResponse, model.LoginResponse, error) {
	var user_response model.GetUserResponse
	var login_response model.LoginResponse
	user, error := authService.userRepository.CheckUsername(request.Username)

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

		expirationTime := time.Now().Add(time.Minute * 60)
		claims := &model.Claims{
			SignatureKey: general.GetMD5Hash(request.Username, strconv.Itoa(user[0].Id)),
			Username:     request.Username,
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

func (authService *authService) Register(request model.RegisterRequest) (entity.User, error) {
	var user entity.User
	date_now := time.Now()

	users, error := authService.userRepository.CheckUsername(request.Username)

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

			_, error = authService.authRepository.Register(request)
			users, error = authService.userRepository.CheckUsername(request.Username)
			user = users[0]
		}
	}

	return user, error
}
