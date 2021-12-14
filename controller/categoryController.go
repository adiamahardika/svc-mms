package controller

import (
	"fmt"
	"net/http"
	"svc-monitoring-maintenance/general"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type categoryController struct {
	categoryService service.CategoryServiceInterface
}

func CategoryController(categoryService service.CategoryServiceInterface) *categoryController{
	return &categoryController{categoryService}
}

func (controller *categoryController) GetCategory(context *gin.Context) {
	category, error := controller.categoryService.GetCategory()
	description := []string{}

	if (error == nil) {
		description = append(description, "Success")

		status := model.StandardResponse{
			HttpStatus: http.StatusOK,
			StatusCode: general.SuccessStatusCode,
			Description: description,
		}
		context.JSON(http.StatusOK, gin.H{
			"status" : status,
			"result" : category,
		})
	} else {
		description = append(description, error.Error())

		status := model.StandardResponse{
			HttpStatus: http.StatusBadRequest,
			StatusCode: general.ErrorStatusCode,
			Description: description,
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"status" : status,
		})
	}
}

func (controller *categoryController) CreateCategory(context *gin.Context) {
	var request model.CreateCategoryRequest

	error := context.ShouldBindJSON(&request)
	description := []string{}

	if (error != nil) {
		for _, value := range error.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", value.Field(), value.ActualTag())
			description = append(description, errorMessage)
		}

		status := model.StandardResponse{
			HttpStatus: http.StatusBadRequest,
			StatusCode: general.ErrorStatusCode,
			Description: description,
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"status" : status,
		})
	}else {
		
		category, error := controller.categoryService.CreateCategory(request)

		if (error == nil) {
			
			description = append(description, "Success")

			status := model.StandardResponse{
				HttpStatus: http.StatusOK,
				StatusCode: general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status" : status,
				"result" : category,
			})

		} else {

			description = append(description, error.Error())

			status := model.StandardResponse{
				HttpStatus: http.StatusBadRequest,
				StatusCode: general.ErrorStatusCode,
				Description: description,
			}
			context.JSON(http.StatusBadRequest, gin.H{
				"status" : status,
			})

		}
	}
}