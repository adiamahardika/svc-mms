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

type taskListController struct {
	taskListService service.TaskListServiceInterface
}

func TaskListController(taskListService service.TaskListServiceInterface) *taskListController {
	return &taskListController{taskListService}
}

func (controller *taskListController) GetTaskList(context *gin.Context) {
	var request model.GetTaskListRequest

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

		task_list, error := controller.taskListService.GetTaskList(&request)

		if (error == nil) {
			
			description = append(description, "Success")

			status := model.StandardResponse{
				HttpStatus: http.StatusOK,
				StatusCode: general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status" : status,
				"result" : task_list,
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

func (controller *taskListController) UpdateTaskListController(context *gin.Context) {
	var request model.UpdateTaskListRequest

	error := context.ShouldBind(&request)
	description := []string{}

	if (error != nil) {

		for _, value := range error.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", value.Field(), value.ActualTag())
			description = append(description, errorMessage)
		}

		fmt.Println(error)
		description = append(description, "Bad Request")

		status := model.StandardResponse{
			HttpStatus: http.StatusBadRequest,
			StatusCode: general.ErrorStatusCode,
			Description: description,
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"status" : status,
		})

	} else {

		_, error := controller.taskListService.UpdateTaskList(request, context)

		if (error == nil) {
			
			description = append(description, "Success")

			status := model.StandardResponse{
				HttpStatus: http.StatusOK,
				StatusCode: general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status" : status,
				// "result" : task_list,
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