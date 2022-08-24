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

type webPermissionController struct {
	webPermissionService service.WebPermissionServiceInterface
	logService           service.LogServiceInterface
}

func WebPermissonController(webPermissionService service.WebPermissionServiceInterface, logService service.LogServiceInterface) *webPermissionController {
	return &webPermissionController{webPermissionService, logService}
}

func (controller *webPermissionController) GetPermission(context *gin.Context) {
	description := []string{}
	http_status := http.StatusOK
	var status *model.StandardResponse

	permission, error := controller.webPermissionService.GetWebPermission()

	if error == nil {
		description = append(description, "Success")

		status = &model.StandardResponse{
			HttpStatus:  http.StatusOK,
			StatusCode:  general.SuccessStatusCode,
			Description: description,
		}
		context.JSON(http.StatusOK, gin.H{
			"status":         status,
			"web_permission": permission,
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
	parse_permission, _ := json.Marshal(permission)
	var result = fmt.Sprintf("{\"status\": %s, \"web_permission\": %s}", string(parse_status), string(parse_permission))
	controller.logService.CreateLog(context, "", result, time.Now(), http_status)
}
