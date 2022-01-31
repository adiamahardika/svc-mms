package service

import (
	"fmt"
	"net/http"
	"os"
	"svc-monitoring-maintenance/general"
	"svc-monitoring-maintenance/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthServiceInterface interface {
	Authentication(context *gin.Context)
}

type authService struct {
}

func AuthService() *authService {
	return &authService{}
}

func Authentication() gin.HandlerFunc {
	return func(context *gin.Context) {

		tokenString := context.Request.Header.Get("token")
		claims := &model.Claims{}
		description := []string{}
		jwtKey := []byte(os.Getenv("API_SECRET"))
		var status model.StandardResponse

		token, error := jwt.ParseWithClaims(tokenString, claims,
			func(t *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

		if error != nil || token == nil {
			error = fmt.Errorf(fmt.Sprintf("Please provide token!"))
		} else if !token.Valid {
			error = fmt.Errorf(fmt.Sprintf("Your token is invalid!"))
		}
		fmt.Println(error)
		if error != nil {
			description = append(description, error.Error())
			status = model.StandardResponse{
				HttpStatus:  http.StatusBadRequest,
				StatusCode:  general.ErrorStatusCode,
				Description: description,
			}
			context.JSON(http.StatusBadRequest, gin.H{
				"status": status,
			})
		} else {
			context.Next()
		}
	}
}
