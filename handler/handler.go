package handler

import (
	controllers "a/src/controller"
	"a/src/database"
	"a/src/repositories"
	"a/src/service"
)

var (
	ticketRepository = repositories.TicketRepository{}
	db               = database.ScyllaConn{}
	ticketService    = service.New(&ticketRepository, &db)
	ticketController = controllers.New(&ticketService)
)
