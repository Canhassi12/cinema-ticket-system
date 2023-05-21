package main

import (
	"log"
	"ticket-system/src/database"
	"ticket-system/src/repositories"
	"ticket-system/src/service"
	"time"
)

func main() {
	ticketRepository := repositories.TicketRepository{}
	db := database.ScyllaConn{}
	ticketService := service.New(&ticketRepository, &db)

	session := db.Conn()

	for {
		if err := ticketService.ReleaseTicket(session); err != nil {
			log.Printf("worker error: %s\n", err.Error())
		}

		time.Sleep(time.Minute + time.Second*30)
	}
}
