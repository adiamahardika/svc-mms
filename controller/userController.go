package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"svc-monitoring-maintenance/general"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userController struct {
	userService service.UserServiceInterface
	logService  service.LogServiceInterface
}

func UserController(userService service.UserServiceInterface, logService service.LogServiceInterface) *userController {
	return &userController{userService, logService}
}

func (controller *userController) GetUser(context *gin.Context) {
	var request model.GetUserRequest

	error := context.ShouldBindJSON(&request)
	description := []string{}
	http_status := http.StatusOK
	var status model.StandardResponse
	var user []model.GetUserResponse

	if error != nil {
		for _, value := range error.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", value.Field(), value.ActualTag())
			description = append(description, errorMessage)
		}
		http_status = http.StatusBadRequest

		status = model.StandardResponse{
			HttpStatus:  http.StatusBadRequest,
			StatusCode:  general.ErrorStatusCode,
			Description: description,
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"status": status,
		})
	} else {

		user, error = controller.userService.GetUser(request)
		if error == nil {

			description = append(description, "Success")

			status = model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
				"result": user,
			})

		} else {

			description = append(description, error.Error())
			http_status = http.StatusBadRequest

			status = model.StandardResponse{
				HttpStatus:  http.StatusBadRequest,
				StatusCode:  general.ErrorStatusCode,
				Description: description,
			}
			context.JSON(http.StatusBadRequest, gin.H{
				"status": status,
			})

		}
	}
	parse_request, _ := json.Marshal(request)
	parse_status, _ := json.Marshal(status)
	parse_user, _ := json.Marshal(user)
	var result = fmt.Sprintf("{\"status\": %s, \"result\": %s}", string(parse_status), string(parse_user))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}

func (controller *userController) ChangePassword(context *gin.Context) {
	var request model.ChangePassRequest

	error := context.ShouldBindJSON(&request)
	description := []string{}
	http_status := http.StatusOK
	var status model.StandardResponse
	var user model.GetUserResponse

	if error != nil {
		for _, value := range error.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", value.Field(), value.ActualTag())
			description = append(description, errorMessage)
		}
		http_status = http.StatusBadRequest

		status = model.StandardResponse{
			HttpStatus:  http.StatusBadRequest,
			StatusCode:  general.ErrorStatusCode,
			Description: description,
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"status": status,
		})
	} else {

		user, error = controller.userService.ChangePassword(request)

		if error == nil {

			description = append(description, "Success")

			status = model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
				"result": user,
			})

		} else {

			description = append(description, error.Error())
			http_status = http.StatusBadRequest

			status = model.StandardResponse{
				HttpStatus:  http.StatusBadRequest,
				StatusCode:  general.ErrorStatusCode,
				Description: description,
			}
			context.JSON(http.StatusBadRequest, gin.H{
				"status": status,
			})

		}
	}
	parse_request, _ := json.Marshal(request)
	parse_status, _ := json.Marshal(status)
	parse_user, _ := json.Marshal(user)
	var result = fmt.Sprintf("{\"status\": %s, \"result\": %s}", string(parse_status), string(parse_user))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}

func (controller *userController) ResetPassword(context *gin.Context) {
	var request model.ResetPassword

	error := context.ShouldBindJSON(&request)
	description := []string{}
	http_status := http.StatusOK
	var status model.StandardResponse
	var user model.GetUserResponse

	if error != nil {
		for _, value := range error.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", value.Field(), value.ActualTag())
			description = append(description, errorMessage)
		}
		http_status = http.StatusBadRequest

		status = model.StandardResponse{
			HttpStatus:  http.StatusBadRequest,
			StatusCode:  general.ErrorStatusCode,
			Description: description,
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"status": status,
		})
	} else {

		user, error = controller.userService.ResetPassword(request)

		if error == nil {

			description = append(description, "Success")

			status = model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
				"result": user,
			})

		} else {

			description = append(description, error.Error())
			http_status = http.StatusBadRequest

			status = model.StandardResponse{
				HttpStatus:  http.StatusBadRequest,
				StatusCode:  general.ErrorStatusCode,
				Description: description,
			}
			context.JSON(http.StatusBadRequest, gin.H{
				"status": status,
			})

		}
	}
	parse_request, _ := json.Marshal(request)
	parse_status, _ := json.Marshal(status)
	parse_user, _ := json.Marshal(user)
	var result = fmt.Sprintf("{\"status\": %s, \"result\": %s}", string(parse_status), string(parse_user))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}

func (controller *userController) GetDetailUser(context *gin.Context) {
	user_id := context.Param("user-id")

	description := []string{}
	http_status := http.StatusOK
	var status model.StandardResponse

	user, error := controller.userService.GetDetailUser(user_id)

	if error == nil {

		description = append(description, "Success")

		status = model.StandardResponse{
			HttpStatus:  http.StatusOK,
			StatusCode:  general.SuccessStatusCode,
			Description: description,
		}
		context.JSON(http.StatusOK, gin.H{
			"status": status,
			"result": user,
		})

	} else {

		description = append(description, error.Error())
		http_status = http.StatusBadRequest

		status = model.StandardResponse{
			HttpStatus:  http.StatusBadRequest,
			StatusCode:  general.ErrorStatusCode,
			Description: description,
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"status": status,
		})

	}
	parse_status, _ := json.Marshal(status)
	parse_user, _ := json.Marshal(user)
	var result = fmt.Sprintf("{\"status\": %s, \"result\": %s}", string(parse_status), string(parse_user))
	controller.logService.CreateLog(context, "", result, time.Now(), http_status)
}

func (controller *userController) UpdateKeyHp(context *gin.Context) {
	var request *model.UpdateKeyHpRequest

	error := context.ShouldBindJSON(&request)
	description := []string{}
	http_status := http.StatusOK
	var status *model.StandardResponse
	var user model.UpdateKeyHpRequest

	if error != nil {
		for _, value := range error.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", value.Field(), value.ActualTag())
			description = append(description, errorMessage)
		}
		http_status = http.StatusBadRequest

		status = &model.StandardResponse{
			HttpStatus:  http.StatusBadRequest,
			StatusCode:  general.ErrorStatusCode,
			Description: description,
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"status": status,
		})
	} else {

		user, error = controller.userService.UpdateKeyHp(request)

		if error == nil {

			description = append(description, "Success")

			status = &model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
				"result": user,
			})

		} else {

			description = append(description, error.Error())
			http_status = http.StatusBadRequest

			status = &model.StandardResponse{
				HttpStatus:  http.StatusBadRequest,
				StatusCode:  general.ErrorStatusCode,
				Description: description,
			}
			context.JSON(http.StatusBadRequest, gin.H{
				"status": status,
			})

		}
	}
	parse_request, _ := json.Marshal(request)
	parse_status, _ := json.Marshal(status)
	parse_user, _ := json.Marshal(user)
	var result = fmt.Sprintf("{\"status\": %s, \"result\": %s}", string(parse_status), string(parse_user))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}
