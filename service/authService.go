package service

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"svc-monitoring-maintenance/general"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type authService struct {
	repository repository.UserRepositoryInterface
}

func AuthService(repository repository.UserRepositoryInterface) *authService {
	return &authService{repository}
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

		user, error := authService.repository.CheckUsername(claims.Username)
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
