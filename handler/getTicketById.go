package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

var (
	getByIdFunc = ticketController.GetById
)

func GetTicketById(w http.ResponseWriter, r *http.Request) {

	ticketModel, err := getByIdFunc(chi.URLParam(r, "ticketId"))

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
}
