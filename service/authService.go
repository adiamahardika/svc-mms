package service

import (
	"fmt"
	"net/http"
	"os"
	"svc-monitoring-maintenance/general"
	"svc-monitoring-maintenance/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(context *gin.Context) {

		tokenString := context.Request.Header.Get("token")
		claims := &model.Claims{}
		description := []string{}
		jwtKey := []byte(os.Getenv("API_SECRET"))
		var status model.StandardResponse

		token, error := jwt.ParseWithClaims(tokenString, claims,
			func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

		if token == nil {
			error = fmt.Errorf(fmt.Sprintf("Please provide token!"))
		} else if error.Error() == "signature is invalid" {
			error = fmt.Errorf(fmt.Sprintf("Your token is invalid!"))
		} else if time.Now().Unix() > claims.ExpiresAt {
			error = fmt.Errorf(fmt.Sprintf("Your token is expired!"))
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
		} else {
			context.Next()
		}
	}
}
