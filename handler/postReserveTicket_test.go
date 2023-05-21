package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"ticket-system/src/requests"
)

func mockReserveTicketSuccess(ticketId string, data requests.Payload) error {
	return nil
}

func TestReserveTicket(t *testing.T) {
	tests := map[string]struct {
		expectStatusCode int
		ticketId         string
		reserveFunc      func(ticketId string, data requests.Payload) error
		body             string
	}{
		"success: reserve ticket": {
			expectStatusCode: http.StatusOK,
			ticketId:         "id001",
			reserveFunc:      mockReserveTicketSuccess,
			body:             `{"price":"half", "seat": 2, "movie": "alou", "userId": "teste123"}`,
		},
		// "failed: empty ticket Id": {
		// 	expectStatusCode: http.StatusBadRequest,
		// 	ticketId:         "",
		// },
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			updateReserveTicketFunc = test.reserveFunc

			w := httptest.NewRecorder()
			r := httptest.NewRequest("PUT", fmt.Sprintf("/reserve/%s", test.ticketId), strings.NewReader(test.body))

			UpdateReserveTicket(w, r)
			response := w.Result()

			if response.StatusCode != test.expectStatusCode {
				t.Errorf(
					"expect code %v, got %v",
					test.expectStatusCode, response.StatusCode,
				)
			}
		})
	}
}
