package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"svc-monitoring-maintenance/general"
	"svc-monitoring-maintenance/model"
	"svc-monitoring-maintenance/service"
)

type ticketController struct {
	ticketService service.TicketServiceInterface
}

func TicketController(ticketService service.TicketServiceInterface) *ticketController {
	return &ticketController{ticketService}
}

func (controller *ticketController) GetAll(context *gin.Context) {

	list_ticket, error := controller.ticketService.GetAll()
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
			"result" : list_ticket,
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

func (controller *ticketController) GetTicket(context *gin.Context) {
	var request model.GetTicketRequest

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

	} else {
		
		list_ticket, error := controller.ticketService.GetTicket(request)

		if (error == nil) {
			
			description = append(description, "Success")

			status := model.StandardResponse{
				HttpStatus: http.StatusOK,
				StatusCode: general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status" : status,
				"result" : list_ticket,
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

func (controller *ticketController) CountTicketByStatus(context *gin.Context) {

	var request model.CountTicketByStatusRequest

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
	} else {
		
		list_status, error := controller.ticketService.CountTicketByStatus(request)

		if (error == nil) {
			
			description = append(description, "Success")

			status := model.StandardResponse{
				HttpStatus: http.StatusOK,
				StatusCode: general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status" : status,
				"result" : list_status,
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

func (controller *ticketController) CreateTicket(context *gin.Context) {
	var request model.CreateTicketRequest

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
	} else {
		
		ticket, error := controller.ticketService.CreateTicket(request)

		if (error == nil) {
			
			description = append(description, "Success")

			status := model.StandardResponse{
				HttpStatus: http.StatusOK,
				StatusCode: general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status" : status,
				"result" : ticket,
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

func (controller *ticketController) AssignTicket(context *gin.Context) {
	var request model.AssignTicketRequest

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
	} else {
		
		ticket, error := controller.ticketService.AssignTicket(request)

		if (error == nil) {
			
			description = append(description, "Success")

			status := model.StandardResponse{
				HttpStatus: http.StatusOK,
				StatusCode: general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status" : status,
				"result" : ticket,
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

func (controller *ticketController) UpdateTicketStatus(context *gin.Context) {
	var request model.UpdateTicketStatusRequest

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
	} else {
		
		ticket, error := controller.ticketService.UpdateTicketStatus(request)

		if (error == nil) {
			
			description = append(description, "Success")

			status := model.StandardResponse{
				HttpStatus: http.StatusOK,
				StatusCode: general.SuccessStatusCode,
				Description: description,
			}
			context.JSON(http.StatusOK, gin.H{
				"status" : status,
				"result" : ticket,
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

func (controller *ticketController) GetDetailTicket(context *gin.Context) {
	request := context.Param("ticket-code")

	description := []string{}
		
	ticket, error := controller.ticketService.GetDetailTicket(request)

	if (error == nil) {
		
		description = append(description, "Success")

		status := model.StandardResponse{
			HttpStatus: http.StatusOK,
			StatusCode: general.SuccessStatusCode,
			Description: description,
		}
		context.JSON(http.StatusOK, gin.H{
			"status" : status,
			"result" : ticket,
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