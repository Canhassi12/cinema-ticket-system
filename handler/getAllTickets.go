package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

var (
	getAllTicketsFunc = ticketController.GetAll
)

func GetAllTickets(w http.ResponseWriter, r *http.Request) {
	ticketModels, err := getAllTicketsFunc(chi.URLParam(r, "movieId"))

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
}
