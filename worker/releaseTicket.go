package main

import (
	"a/src/database"
	"a/src/repositories"
	"a/src/service"
	"log"
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
