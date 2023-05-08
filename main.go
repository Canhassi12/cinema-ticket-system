package main

import (
	"a/src/database"
	"a/src/repositories"
	"a/src/requests"
	"a/src/service"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	controllers "a/src/controller"
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
			return
		}

		jsonBytes, err := json.Marshal(ticketModel)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte(jsonBytes))
	})

	r.Get("/tickets/{movieId}", func(w http.ResponseWriter, r *http.Request) {
		ticketModels, err := ticketController.GetAll(chi.URLParam(r, "movieId"))

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		jsonBytes, err := json.Marshal(ticketModels)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte(jsonBytes))
	})

	r.Post("/reserve/{ticketId}", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		var m requests.Payload = requests.Payload{}
		err = json.Unmarshal(body, &m)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Error unmarshalling request body"))
			return
		}

		if err := ticketController.ReserveForPay(chi.URLParam(r, "ticketId"), m); err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte("you reserve this seat, you have ten minutes to pay that"))
	})

	r.Put("/reserve/{ticketId}", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		var m requests.Payload = requests.Payload{}
		err = json.Unmarshal(body, &m)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Error unmarshalling request body"))
			return
		}

		if err := ticketController.Pay(chi.URLParam(r, "ticketId"), m); err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte("the payment was successful"))
	})

	http.ListenAndServe(":3000", r)
}
