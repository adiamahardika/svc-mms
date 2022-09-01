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

type hwReplacementStatusController struct {
	hwReplacementStatusService service.HwReplacementStatusServiceInterface
	logService                 service.LogServiceInterface
}

func HwReplacementStatusController(hwReplacementStatusService service.HwReplacementStatusServiceInterface, logService service.LogServiceInterface) *hwReplacementStatusController {
	return &hwReplacementStatusController{hwReplacementStatusService, logService}
}

func (controller *hwReplacementStatusController) GetHwReplacementStatus(context *gin.Context) {
	description := []string{}
	http_status := http.StatusOK
	var status *model.StandardResponse

	hwReplacementstatus, error := controller.hwReplacementStatusService.GetHwReplacementStatus()

	if error == nil {
		description = append(description, "Success")

		status = &model.StandardResponse{
			HttpStatus:  http.StatusOK,
			StatusCode:  general.SuccessStatusCode,
			Description: description,
		}
		context.JSON(http.StatusOK, gin.H{
			"status":  status,
			"content": hwReplacementstatus,
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
	parse_content, _ := json.Marshal(hwReplacementstatus)
	var result = fmt.Sprintf("{\"status\": %s, \"content\": %s}", string(parse_status), string(parse_content))
	controller.logService.CreateLog(context, "", result, time.Now(), http_status)
}

func (controller *hwReplacementStatusController) CreateHwReplacementStatus(context *gin.Context) {
	var request *entity.HwReplacementStatus

	error := context.ShouldBindJSON(&request)
	description := []string{}
	http_status := http.StatusOK
	var status model.StandardResponse
	var hw_replacement_status []entity.HwReplacementStatus

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

		hw_replacement_status, error = controller.hwReplacementStatusService.CreateHwReplacementStatus(request)

		if error == nil {

			description = append(description, "Success")

			status = model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
				"result": hw_replacement_status,
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
	parse_hw_replacement_status, _ := json.Marshal(hw_replacement_status)
	var result = fmt.Sprintf("{\"status\": %s, \"result\": %s}", string(parse_status), string(parse_hw_replacement_status))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}

func (controller *hwReplacementStatusController) UpdateHwReplacementStatus(context *gin.Context) {
	var request *entity.HwReplacementStatus

	error := context.ShouldBindJSON(&request)
	description := []string{}
	http_status := http.StatusOK
	var status model.StandardResponse
	var hw_replacement_status []entity.HwReplacementStatus

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

		hw_replacement_status, error = controller.hwReplacementStatusService.UpdateHwReplacementStatus(request)

		if error == nil {

			description = append(description, "Success")

			status = model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
				"result": hw_replacement_status,
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
	parse_hw_replacement_status, _ := json.Marshal(hw_replacement_status)
	var result = fmt.Sprintf("{\"status\": %s, \"result\": %s}", string(parse_status), string(parse_hw_replacement_status))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}
