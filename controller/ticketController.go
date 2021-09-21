package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"svc-ticket-monitoring/general"
	"svc-ticket-monitoring/model"
	"svc-ticket-monitoring/service"
)

type ticketController struct {
	ticketService service.TicketServiceInterface
}

func TicketController(ticketService service.TicketServiceInterface) *ticketController {
	return &ticketController{ticketService}
}

func (controller *ticketController) FindAll(context *gin.Context) {

	list_ticket, error := controller.ticketService.FindAll()
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

func (controller *ticketController) FindTicket(context *gin.Context) {
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
		
		list_ticket, error := controller.ticketService.FindTicket(request)

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