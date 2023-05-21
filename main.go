package main

import (
	"net/http"
	"ticket-system/handler"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ticket/{ticketId}", handler.GetTicketById)

	r.Get("/tickets/{movieId}", handler.GetAllTickets)

	r.Post("/reserve/{ticketId}", handler.PostReserveTicket)

	r.Put("/reserve/{ticketId}", handler.UpdateReserveTicket)

	http.ListenAndServe(":3000", r)
}
