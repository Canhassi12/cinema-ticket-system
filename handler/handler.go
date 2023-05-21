package handler

import (
	controllers "ticket-system/src/controller"
	"ticket-system/src/database"
	"ticket-system/src/repositories"
	"ticket-system/src/service"
)

var (
	ticketRepository = repositories.TicketRepository{}
	db               = database.ScyllaConn{}
	ticketService    = service.New(&ticketRepository, &db)
	ticketController = controllers.New(&ticketService)
)
