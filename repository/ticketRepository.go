package repository

import (
	"fmt"
	"svc-ticket-monitoring/entity"
	"svc-ticket-monitoring/model"

	"gorm.io/gorm"
)
type TicketRepositoryInterface interface {
	FindAll() 									([]entity.Ticket, error)
	FindTicket(request model.GetTicketRequest)	([]entity.Ticket, error)
}

type repository struct {
	db *gorm.DB
}

func TicketRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repo *repository) FindAll() ([]entity.Ticket, error) {
	var ticket []entity.Ticket

	error := repo.db.Raw("SELECT * FROM ticket ORDER BY tgl_dibuat ASC").Scan(&ticket).Error
	fmt.Println(ticket) 
	return ticket, error
}

func (repo *repository) FindTicket(request model.GetTicketRequest) ([]entity.Ticket, error) {
	var ticket []entity.Ticket
	
	error := repo.db.Raw("SELECT * FROM (SELECT * FROM ticket WHERE prioritas LIKE @Priority AND status LIKE @Status AND assigned_to LIKE @AssignedTo AND assigned_to_team LIKE @AssignedToTeam AND username_pembuat LIKE @UsernamePembuat AND tgl_dibuat >= @StartDate AND tgl_dibuat <= @EndDate ORDER BY tgl_diperbarui DESC) as tbl WHERE judul LIKE @Search OR kode_ticket LIKE @Search OR lokasi LIKE @Search OR terminal_id LIKE @Search OR email LIKE @Search", model.GetTicketRequest{
		PageNo:          request.PageNo,
		PageSize:        request.PageSize,
		SortBy:          "%" + request.SortBy + "%",
		Search:          "%" + request.Search + "%",
		Status:          "%" + request.Status + "%",
		Priority:        "%" + request.Priority + "%",
		AssignedTo:      "%" + request.AssignedTo + "%",
		AssignedToTeam:  "%" + request.AssignedToTeam + "%",
		UsernamePembuat: "%" + request.UsernamePembuat + "%",
		StartDate: 		request.StartDate,
		EndDate: 		request.EndDate,
		}).Find(&ticket).Error
	
	
	// subQuery := repo.db.Select("ticket.*").Table("ticket").Where("prioritas LIKE @Priority AND status LIKE @Status AND assigned_to LIKE @AssignedTo AND assigned_to_team LIKE @AssignedToTeam AND username_pembuat LIKE @UsernamePembuat ORDER BY tgl_diperbarui DESC", model.GetTicketRequest{
	// 	SortBy:          "%" + request.SortBy + "%",
	// 	Status:          "%" + request.Status + "%",
	// 	Priority:        "%" + request.Priority + "%",
	// 	AssignedTo:      "%" + request.AssignedTo + "%",
	// 	AssignedToTeam:  "%" + request.AssignedToTeam + "%",
	// 	UsernamePembuat: "%" + request.UsernamePembuat + "%",
	// 	})
	// error := repo.db.Raw("SELECT * FROM (?) as new_table WHERE judul LIKE @Search OR kode_ticket LIKE @Search OR lokasi LIKE @Search OR terminal_id LIKE @Search OR email LIKE @Search", subQuery, model.GetTicketRequest{
	// Search:          "%" + request.Search + "%",
	// }).Find(&ticket).Error
	fmt.Println( request.AssignedToTeam)

	// error := repo.db.Raw(`SELECT * FROM ticket WHERE judul LIKE @Search`, model.GetTicketRequest{
	// 		PageNo:          request.PageNo,
	// 		PageSize:        request.PageSize,
	// 		SortBy:          request.SortBy,
	// 		Search:          "%" + request.Search + "%",
	// 		Status:          request.Status,
	// 		Priority:        request.Priority,
	// 		AssignedTo:      request.AssignedTo,
	// 		AssignedToTeam:  request.AssignedToTeam,
	// 		UsernamePembuat: request.UsernamePembuat,
	// 		}).Find(&ticket).Error
		
		return ticket, error
	}