package controller

import (
	"fmt"
	"net/http"
	"svc-monitoring-maintenance/general"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userController struct {
	userService service.UserServiceInterface
}

func UserController(userService service.UserServiceInterface) *userController {
	return &userController{userService}
}

func (controller *userController) GetUser(context *gin.Context) {
	var request model.GetUserRequest

	error := context.ShouldBindJSON(&request)
	description := []string{}

	if (error != nil) {
		for _, value := range error.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", value.Field(), value.ActualTag())
			description = append(description, errorMessage)
		}

		status := model.StandardResponse{
			HttpStatus: http.StatusBadRequest,
			StatusCode: general.ErrorStatusCode,
			Description: description,
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"status" : status,
		})
	} else {

		user, error := controller.userService.GetUser(request)
		if (error == nil) {
			
			description = append(description, "Success")

			status := model.StandardResponse{
				HttpStatus: http.StatusOK,
				StatusCode: general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status" : status,
				"result" : user,
			})

		} else {

			description = append(description, error.Error())

			status := model.StandardResponse{
				HttpStatus: http.StatusBadRequest,
				StatusCode: general.ErrorStatusCode,
				Description: description,
			}
			context.JSON(http.StatusBadRequest, gin.H{
				"status" : status,
			})

		}
	}
}

func (controller *userController) Login(context *gin.Context) {
	var request model.LoginRequest

	error := context.ShouldBindJSON(&request)
	description := []string{}

	if (error != nil) {
		for _, value := range error.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", value.Field(), value.ActualTag())
			description = append(description, errorMessage)
		}

		status := model.StandardResponse{
			HttpStatus: http.StatusBadRequest,
			StatusCode: general.ErrorStatusCode,
			Description: description,
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"status" : status,
		})
	} else {
		
		user, error := controller.userService.Login(request)

		if (error == nil) {
			
			description = append(description, "Success")

			status := model.StandardResponse{
				HttpStatus: http.StatusOK,
				StatusCode: general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status" : status,
				"result" : user,
			})

		} else {

			description = append(description, error.Error())

			status := model.StandardResponse{
				HttpStatus: http.StatusOK,
				StatusCode: general.ErrorStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status" : status,
			})

		}
	}
}

func (controller *userController) ChangePassword(context *gin.Context) {
	var request model.ChangePassRequest

	error := context.ShouldBindJSON(&request)
	description := []string{}

	if (error != nil) {
		for _, value := range error.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", value.Field(), value.ActualTag())
			description = append(description, errorMessage)
		}

		status := model.StandardResponse{
			HttpStatus: http.StatusBadRequest,
			StatusCode: general.ErrorStatusCode,
			Description: description,
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"status" : status,
		})
	} else {
		
		user, error := controller.userService.ChangePassword(request)

		if (error == nil) {
			
			description = append(description, "Success")

			status := model.StandardResponse{
				HttpStatus: http.StatusOK,
				StatusCode: general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status" : status,
				"result" : user,
			})

		} else {

			description = append(description, error.Error())

			status := model.StandardResponse{
				HttpStatus: http.StatusBadRequest,
				StatusCode: general.ErrorStatusCode,
				Description: description,
			}
			context.JSON(http.StatusBadRequest, gin.H{
				"status" : status,
			})

		}
	}
}

func (controller *userController) ResetPassword(context *gin.Context) {
	var request model.ResetPassword

	error := context.ShouldBindJSON(&request)
	description := []string{}

	if (error != nil) {
		for _, value := range error.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", value.Field(), value.ActualTag())
			description = append(description, errorMessage)
		}

		status := model.StandardResponse{
			HttpStatus: http.StatusBadRequest,
			StatusCode: general.ErrorStatusCode,
			Description: description,
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"status" : status,
		})
	} else {
		
		user, error := controller.userService.ResetPassword(request)

		if (error == nil) {
			
			description = append(description, "Success")

			status := model.StandardResponse{
				HttpStatus: http.StatusOK,
				StatusCode: general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status" : status,
				"result" : user,
			})

		} else {

			description = append(description, error.Error())

			status := model.StandardResponse{
				HttpStatus: http.StatusBadRequest,
				StatusCode: general.ErrorStatusCode,
				Description: description,
			}
			context.JSON(http.StatusBadRequest, gin.H{
				"status" : status,
			})

		}
	}
}

func (controller *userController) Register(context *gin.Context) {
	var request model.RegisterRequest

	error := context.ShouldBindJSON(&request)
	description := []string{}

	if (error != nil) {
		for _, value := range error.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", value.Field(), value.ActualTag())
			description = append(description, errorMessage)
		}

		status := model.StandardResponse{
			HttpStatus: http.StatusBadRequest,
			StatusCode: general.ErrorStatusCode,
			Description: description,
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"status" : status,
		})
	} else {
		
		user, error := controller.userService.Register(request)

		if (error == nil) {
			
			description = append(description, "Success")

			status := model.StandardResponse{
				HttpStatus: http.StatusOK,
				StatusCode: general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status" : status,
				"result" : user,
			})

		} else {

			description = append(description, error.Error())

			status := model.StandardResponse{
				HttpStatus: http.StatusBadRequest,
				StatusCode: general.ErrorStatusCode,
				Description: description,
			}
			context.JSON(http.StatusBadRequest, gin.H{
				"status" : status,
			})

		}
	}
}