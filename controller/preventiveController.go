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

type preventiveController struct {
	preventiveService service.PreventiveServiceInterface
	logService        service.LogServiceInterface
}

func PreventiveController(preventiveService service.PreventiveServiceInterface, logService service.LogServiceInterface) *preventiveController {
	return &preventiveController{preventiveService, logService}
}

func (controller *preventiveController) CreatePreventive(context *gin.Context) {
	var request []entity.Preventive

	error := context.ShouldBindJSON(&request)
	description := []string{}
	http_status := http.StatusOK
	var status model.StandardResponse
	var list_preventive []entity.Preventive

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

		list_preventive, error = controller.preventiveService.CreatePreventive(request)

		if error == nil {
			description = append(description, "Success")

			status = model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
				"result": list_preventive,
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
	parse_preventive, _ := json.Marshal(list_preventive)
	var result = fmt.Sprintf("{\"status\": %s, \"result\": %s}", string(parse_status), string(parse_preventive))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}

func (controller *preventiveController) GetPreventive(context *gin.Context) {
	var request *model.GetPreventiveRequest

	error := context.ShouldBindJSON(&request)
	description := []string{}
	http_status := http.StatusOK
	var status model.StandardResponse
	var list_preventive []model.GetGroupPreventiveResponse
	var total_pages int

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

		list_preventive, total_pages, error = controller.preventiveService.GetPreventive(request)

		if error == nil {
			description = append(description, "Success")

			status = model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status":      status,
				"result":      list_preventive,
				"page":        request.PageNo,
				"total_pages": total_pages,
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
	parse_preventive, _ := json.Marshal(list_preventive)
	var result = fmt.Sprintf("{\"status\": %s, \"result\": %s, \"page\": %d, \"total_pages\": %d}", string(parse_status), string(parse_preventive), request.PageNo, total_pages)
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}

func (controller *preventiveController) UpdatePreventive(context *gin.Context) {
	var request *entity.Preventive

	error := context.ShouldBindJSON(&request)
	description := []string{}
	http_status := http.StatusOK
	var status model.StandardResponse
	var list_preventive entity.Preventive

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

		list_preventive, error = controller.preventiveService.UpdatePreventive(request)

		if error == nil {
			description = append(description, "Success")

			status = model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
				"result": list_preventive,
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
	parse_preventive, _ := json.Marshal(list_preventive)
	var result = fmt.Sprintf("{\"status\": %s, \"result\": %s}", string(parse_status), string(parse_preventive))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}

func (controller *preventiveController) GetDetailPreventive(context *gin.Context) {
	request := context.Param("prev-code")

	description := []string{}
	http_status := http.StatusOK
	var status model.StandardResponse

	ticket, error := controller.preventiveService.GetDetailPreventive(request)

	if error == nil {

		description = append(description, "Success")

		status = model.StandardResponse{
			HttpStatus:  http.StatusOK,
			StatusCode:  general.SuccessStatusCode,
			Description: description,
		}
		context.JSON(http.StatusOK, gin.H{
			"status": status,
			"result": ticket,
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
	parse_ticket, _ := json.Marshal(ticket)
	var result = fmt.Sprintf("{\"status\": %s, \"result\": %s}", string(parse_status), string(parse_ticket))
	controller.logService.CreateLog(context, "", result, time.Now(), http_status)
}

func (controller *preventiveController) CountPreventiveByStatus(context *gin.Context) {
	var request model.CountPreventiveByStatusRequest

	error := context.ShouldBindJSON(&request)
	description := []string{}
	http_status := http.StatusOK
	var status model.StandardResponse
	var list_preventive []model.CountPreventiveByStatusResponse

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

		list_preventive, error = controller.preventiveService.CountPreventiveByStatus(request)

		if error == nil {
			description = append(description, "Success")

			status = model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
				"result": list_preventive,
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
	parse_preventive, _ := json.Marshal(list_preventive)
	var result = fmt.Sprintf("{\"status\": %s, \"result\": %s}", string(parse_status), string(parse_preventive))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}
