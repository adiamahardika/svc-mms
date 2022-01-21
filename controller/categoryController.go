package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"svc-monitoring-maintenance/entity"
	"svc-monitoring-maintenance/general"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type categoryController struct {
	categoryService service.CategoryServiceInterface
	logService      service.LogServiceInterface
}

func CategoryController(categoryService service.CategoryServiceInterface, logService service.LogServiceInterface) *categoryController {
	return &categoryController{categoryService, logService}
}

func (controller *categoryController) GetCategory(context *gin.Context) {
	var request model.GetCategoryRequest

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
		category, error := controller.categoryService.GetCategory(request)
		description := []string{}
		parse_request, _ := json.Marshal(request)
		parse_response, _ := json.Marshal(category)

		if error == nil {
			controller.logService.CreateLog(context, string(parse_request), string(parse_response), time.Now(), http.StatusOK)
			description = append(description, "Success")

			status := model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
				"result": category,
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

func (controller *categoryController) CreateCategory(context *gin.Context) {

	var request model.CreateCategoryRequest

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

		category, error := controller.categoryService.CreateCategory(request)

		if error == nil {

			description = append(description, "Success")

			status := model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
				"result": category,
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

func (controller *categoryController) UpdateCategory(context *gin.Context) {
	var request entity.Category

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

		category, error := controller.categoryService.UpdateCategory(request)

		if error == nil {

			description = append(description, "Success")

			status := model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status": status,
				"result": category,
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

func (controller *categoryController) DeleteCategory(context *gin.Context) {

	id, error := strconv.Atoi(context.Param("category-id"))

	description := []string{}

	error = controller.categoryService.DeleteCategory(id)

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
