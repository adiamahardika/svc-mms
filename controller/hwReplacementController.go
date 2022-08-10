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

type hwReplacementController struct {
	hwReplacementService service.HwReplacementServiceInterface
	logService           service.LogServiceInterface
}

func HwReplacementController(hwReplacementService service.HwReplacementServiceInterface, logService service.LogServiceInterface) *hwReplacementController {
	return &hwReplacementController{hwReplacementService, logService}
}

func (controller *hwReplacementController) CreateHwReplacementController(context *gin.Context) {

	var request *model.HwReplacementRequest

	error := context.ShouldBind(&request)
	description := []string{}
	http_status := http.StatusOK
	var status *model.StandardResponse

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

		_, error := controller.hwReplacementService.CreateHwReplacement(request, context)

		if error == nil {

			description = append(description, "Success")

			status = &model.StandardResponse{
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
	var result = fmt.Sprintf("{\"status\": %s}", string(parse_status))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}
