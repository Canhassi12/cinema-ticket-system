package handler

import (
	"a/src/requests"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

var (
	reserveTicketFunc = ticketController.ReserveForPay
)

func PostReserveTicket(w http.ResponseWriter, r *http.Request) {
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

	if err := reserveTicketFunc(chi.URLParam(r, "ticketId"), m); err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("you reserve this seat, you have ten minutes to pay that"))
}
