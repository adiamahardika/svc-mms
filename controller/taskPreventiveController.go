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

type taskPreventiveController struct {
	taskPreventive service.TaskPreventiveServiceInterface
}

func TaskPreventiveController(taskPreventiveService service.TaskPreventiveServiceInterface) *taskPreventiveController {
	return &taskPreventiveController{taskPreventiveService}
}

func (controller *taskPreventiveController) UpdateTaskPreventiveController(context *gin.Context) {
	var request model.UpdateTaskPreventiveRequest

	error := context.ShouldBind(&request)
	description := []string{}

	if error != nil {

		for _, value := range error.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", value.Field(), value.ActualTag())
			description = append(description, errorMessage)
		}

		status := model.StandardResponse{
			HttpStatus:  http.StatusBadRequest,
			StatusCode:  general.ErrorStatusCode,
			Description: description,
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"status": status,
		})
	} else {

		_, error := controller.taskPreventive.UpdateTaskPreventive(request, context)

		if error == nil {

			description = append(description, "Success")

			status := model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
			})

		} else {

			description = append(description, error.Error())

			status := model.StandardResponse{
				HttpStatus:  http.StatusBadRequest,
				StatusCode:  general.ErrorStatusCode,
				Description: description,
			}
			context.JSON(http.StatusBadRequest, gin.H{
				"status": status,
			})

		}
	}
}

func (controller *taskPreventiveController) GetTaskPreventive(context *gin.Context) {
	var request model.GetTaskPreventiveRequest

	error := context.ShouldBindJSON(&request)
	description := []string{}

	if error != nil {
		for _, value := range error.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", value.Field(), value.ActualTag())
			description = append(description, errorMessage)
		}

		status := model.StandardResponse{
			HttpStatus:  http.StatusBadRequest,
			StatusCode:  general.ErrorStatusCode,
			Description: description,
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"status": status,
		})
	} else {

		task_preventive, error := controller.taskPreventive.GetTaskPreventive(&request)

		if error == nil {

			description = append(description, "Success")

			status := model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
				"result": task_preventive,
			})

		} else {

			description = append(description, error.Error())

			status := model.StandardResponse{
				HttpStatus:  http.StatusBadRequest,
				StatusCode:  general.ErrorStatusCode,
				Description: description,
			}
			context.JSON(http.StatusBadRequest, gin.H{
				"status": status,
			})

		}
	}
}
