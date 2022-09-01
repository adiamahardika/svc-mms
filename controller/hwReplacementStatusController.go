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
