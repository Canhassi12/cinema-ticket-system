package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"ticket-system/src/requests"

	"github.com/go-chi/chi"
)

var (
	updateReserveTicketFunc = ticketController.ReserveForPay
)

func UpdateReserveTicket(w http.ResponseWriter, r *http.Request) {

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

	if err := updateReserveTicketFunc(chi.URLParam(r, "ticketId"), m); err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("the payment was successful"))
}
