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

type taskPreventiveController struct {
	taskPreventive service.TaskPreventiveServiceInterface
	logService     service.LogServiceInterface
}

func TaskPreventiveController(taskPreventiveService service.TaskPreventiveServiceInterface, logService service.LogServiceInterface) *taskPreventiveController {
	return &taskPreventiveController{taskPreventiveService, logService}
}

func (controller *taskPreventiveController) UpdateTaskPreventiveController(context *gin.Context) {
	var request model.UpdateTaskPreventiveRequest

	error := context.ShouldBind(&request)
	description := []string{}
	http_status := http.StatusOK
	var status model.StandardResponse

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

		_, error := controller.taskPreventive.UpdateTaskPreventive(request, context)

		if error == nil {

			description = append(description, "Success")

			status = model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
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
	var result = fmt.Sprintf("{\"status\": %s}", string(parse_status))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}

func (controller *taskPreventiveController) GetTaskPreventive(context *gin.Context) {
	var request model.GetTaskPreventiveRequest

	error := context.ShouldBindJSON(&request)
	description := []string{}
	http_status := http.StatusOK
	var status model.StandardResponse
	var task_preventive []model.GetTaskPreventiveResponse

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

		task_preventive, error = controller.taskPreventive.GetTaskPreventive(&request)

		if error == nil {

			description = append(description, "Success")

			status = model.StandardResponse{
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
	parse_task_preventive, _ := json.Marshal(task_preventive)
	var result = fmt.Sprintf("{\"status\": %s, \"result\": %s}", string(parse_status), string(parse_task_preventive))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}
