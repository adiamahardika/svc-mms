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

type teamController struct {
	teamService service.TeamServiceInterface
}

func TeamController(teamService service.TeamServiceInterface) *teamController {
	return &teamController{teamService}
}

func (controller *teamController) GetAll(context *gin.Context) {

	team, error := controller.teamService.GetTeam()
	description := []string{}

	if error == nil {
		description = append(description, "Success")

		status := model.StandardResponse{
			HttpStatus:  http.StatusOK,
			StatusCode:  general.SuccessStatusCode,
			Description: description,
		}
		context.JSON(http.StatusOK, gin.H{
			"status": status,
			"result": team,
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

func (controller *teamController) CreateTeam(context *gin.Context) {
	var request model.CreateTeamRequest

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

		team, error := controller.teamService.CreateTeam(request)

		if error == nil {

			description = append(description, "Success")

			status := model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
				"result": team,
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

func (controller *teamController) UpdateTeam(context *gin.Context) {
	var request entity.Team

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

		team, error := controller.teamService.UpdateTeam(request)

		if error == nil {

			description = append(description, "Success")

			status := model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
				"result": team,
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

func (controller *teamController) DeleteTeam(context *gin.Context) {

	id, error := strconv.Atoi(context.Param("team-id"))

	description := []string{}

	error = controller.teamService.DeleteService(id)

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
