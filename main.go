package main

import (
	"a/src/database"
	"a/src/repositories"
	"a/src/service"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	controllers "a/src/controller"
	"a/src/requests"
)

func main() {
	ticketRepository := repositories.TicketRepository{}
	db := database.ScyllaConn{}
	ticketService := service.New(&ticketRepository, &db)
	ticketController := controllers.New(&ticketService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/ticket/{ticketId}", func(w http.ResponseWriter, r *http.Request) {
		ticketModel, err := ticketController.GetById(chi.URLParam(r, "ticketId"))

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}

		ticket, err := json.Marshal(ticketModel)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}

		w.Write([]byte(ticket))
	})

	r.Post("/reserve/{userId}", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}

		var m requests.Payload = requests.Payload{}
		err = json.Unmarshal(body, &m)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Error unmarshalling request body"))
		}

		if err := ticketController.ReserveForPay(chi.URLParam(r, "userId"), m); err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}
	})

	http.ListenAndServe(":3000", r)
}
