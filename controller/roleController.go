package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/general"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type roleController struct {
	roleService service.RoleServiceInterface
}

func RoleController(roleService service.RoleServiceInterface) *roleController {
	return &roleController{roleService}
}

func (controller *roleController) GetAll(context *gin.Context) {
	var request model.GetRoleRequest

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

		role, error := controller.roleService.GetRole(request)

		if error == nil {
			description = append(description, "Success")

			status := model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
				"result": role,
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

func (controller *roleController) CreateRole(context *gin.Context) {
	var request model.CreateRoleRequest

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
		role, error := controller.roleService.CreateRole(request)

		if error == nil {
			description = append(description, "Success")

			status := model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
				"result": role,
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

func (controller *roleController) UpdateRole(context *gin.Context) {
	var request entity.Role

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
		role, error := controller.roleService.UpdateRole(request)

		if error == nil {
			description = append(description, "Success")

			status := model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
				"result": role,
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

func (controller *roleController) DeleteRole(context *gin.Context) {

	id, error := strconv.Atoi(context.Param("role-id"))
	description := []string{}

	error = controller.roleService.DeleteRole(id)

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
