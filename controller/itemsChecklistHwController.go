package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/general"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type itemsChecklistHwController struct {
	itemsChecklistHwService service.ItemsChecklistHwServiceInterface
	logService              service.LogServiceInterface
}

func ItemsChecklistHwController(itemsChecklistHwService service.ItemsChecklistHwServiceInterface, logService service.LogServiceInterface) *itemsChecklistHwController {
	return &itemsChecklistHwController{itemsChecklistHwService, logService}
}

func (controller *itemsChecklistHwController) CreateItemsChecklistHw(context *gin.Context) {

	var request *entity.ItemsChecklistHw
	error := context.ShouldBindJSON(&request)
	description := []string{}
	http_status := http.StatusOK
	var status *model.StandardResponse
	var response []entity.ItemsChecklistHw

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

		response, error = controller.itemsChecklistHwService.CreateItemsChecklistHw(request)

		if error == nil {

			description = append(description, "Success")

			status = &model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
				"result": response,
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
	parse_result, _ := json.Marshal(response)
	var result = fmt.Sprintf("{\"status\": %s, \"result\": %s}", string(parse_status), string(parse_result))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}

func (controller *itemsChecklistHwController) GetItemsChecklistHw(context *gin.Context) {

	description := []string{}
	http_status := http.StatusOK
	var status *model.StandardResponse

	response, error := controller.itemsChecklistHwService.GetItemsChecklistHw()

	if error == nil {

		description = append(description, "Success")

		status = &model.StandardResponse{
			HttpStatus:  http.StatusOK,
			StatusCode:  general.SuccessStatusCode,
			Description: description,
		}
		context.JSON(http.StatusOK, gin.H{
			"status": status,
			"result": response,
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

	parse_status, _ := json.Marshal(status)
	parse_result, _ := json.Marshal(response)
	var result = fmt.Sprintf("{\"status\": %s, \"result\": %s}", string(parse_status), string(parse_result))
	controller.logService.CreateLog(context, "", result, time.Now(), http_status)
}

func (controller *itemsChecklistHwController) UpdateItemsChecklistHw(context *gin.Context) {

	var request *entity.ItemsChecklistHw
	error := context.ShouldBindJSON(&request)
	description := []string{}
	http_status := http.StatusOK
	var status *model.StandardResponse
	var response entity.ItemsChecklistHw

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

		response, error = controller.itemsChecklistHwService.UpdateItemsChecklistHw(request)

		if error == nil {

			description = append(description, "Success")

			status = &model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
				"result": response,
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
	parse_result, _ := json.Marshal(response)
	var result = fmt.Sprintf("{\"status\": %s, \"result\": %s}", string(parse_status), string(parse_result))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}
