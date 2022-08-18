package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

	var request *model.GetCategoryRequest

	error := context.ShouldBindJSON(&request)
	description := []string{}
	http_status := http.StatusOK
	var status *model.StandardResponse
	var category []model.CreateCategoryRequest
	var total_pages float64

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
		category, total_pages, error = controller.categoryService.GetCategory(request)

		if error == nil {

			description = append(description, "Success")

			status = &model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status":     status,
				"content":    category,
				"page":       request.PageNo,
				"totalPages": total_pages,
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
	parse_category, _ := json.Marshal(category)
	var result = fmt.Sprintf("{\"status\": %s, \"content\": %s, \"page\": %d, \"totalPages\": %d}", string(parse_status), string(parse_category), request.PageNo, int(total_pages))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}

func (controller *categoryController) CreateCategory(context *gin.Context) {

	var request *model.CreateCategoryRequest

	error := context.ShouldBindJSON(&request)
	description := []string{}
	http_status := http.StatusOK
	var status *model.StandardResponse
	var category model.CreateCategoryRequest

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

		category, error = controller.categoryService.CreateCategory(request)

		if error == nil {

			description = append(description, "Success")

			status = &model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status":  status,
				"content": category,
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
	parse_category, _ := json.Marshal(category)
	var result = fmt.Sprintf("{\"status\": %s, \"content\": %s}", string(parse_status), string(parse_category))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}

func (controller *categoryController) UpdateCategory(context *gin.Context) {
	var request *model.CreateCategoryRequest

	error := context.ShouldBindJSON(&request)
	description := []string{}
	http_status := http.StatusOK
	var status *model.StandardResponse
	var category model.CreateCategoryRequest

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

		category, error = controller.categoryService.UpdateCategory(request)

		if error == nil {

			description = append(description, "Success")

			status = &model.StandardResponse{
				HttpStatus:  http.StatusOK,
				StatusCode:  general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status":  status,
				"content": category,
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
	parse_category, _ := json.Marshal(category)
	var result = fmt.Sprintf("{\"status\": %s, \"content\": %s}", string(parse_status), string(parse_category))
	controller.logService.CreateLog(context, string(parse_request), result, time.Now(), http_status)
}

func (controller *categoryController) DeleteCategory(context *gin.Context) {

	id, error := strconv.Atoi(context.Param("category-id"))

	description := []string{}
	http_status := http.StatusOK
	var status *model.StandardResponse

	error = controller.categoryService.DeleteCategory(&id)

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
	parse_status, _ := json.Marshal(status)
	var result = fmt.Sprintf("{\"status\": %s}", string(parse_status))
	controller.logService.CreateLog(context, "", result, time.Now(), http_status)
}

func (controller *categoryController) GetDetailCategory(context *gin.Context) {

	id := context.Param("id")

	description := []string{}
	http_status := http.StatusOK
	var status *model.StandardResponse

	category, error := controller.categoryService.GetDetailCategory(&id)

	if error == nil {

		description = append(description, "Success")

		status = &model.StandardResponse{
			HttpStatus:  http.StatusOK,
			StatusCode:  general.SuccessStatusCode,
			Description: description,
		}
		context.JSON(http.StatusOK, gin.H{
			"status":  status,
			"content": category,
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
	parse_category, _ := json.Marshal(category)
	var result = fmt.Sprintf("{\"status\": %s, \"content\": %s}", string(parse_status), string(parse_category))
	controller.logService.CreateLog(context, "", result, time.Now(), http_status)
	controller.logService.CreateLog(context, "", result, time.Now(), http_status)
}
