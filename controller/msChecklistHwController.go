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

type msChecklistHwController struct {
	msChecklistHwService service.MsChecklistHwServiceInterface
	logService           service.LogServiceInterface
}

func MsChecklistHwController(msChecklistHwService service.MsChecklistHwServiceInterface, logService service.LogServiceInterface) *msChecklistHwController {
	return &msChecklistHwController{msChecklistHwService, logService}
}

func (controller *msChecklistHwController) CreateMsChecklistHw(context *gin.Context) {

	var request *entity.MsChecklistHw
	error := context.ShouldBindJSON(&request)
	description := []string{}
	http_status := http.StatusOK
	var status *model.StandardResponse
	var response []entity.MsChecklistHw

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

		response, error = controller.msChecklistHwService.CreateMsChecklistHw(request)

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