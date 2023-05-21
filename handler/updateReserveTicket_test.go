package handler

import (
	"a/src/requests"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func mockUpdateReserveTicketSuccess(ticketId string, data requests.Payload) error {
	return nil
}

func TestUpdateReserveTicket(t *testing.T) {
	tests := map[string]struct {
		expectStatusCode int
		ticketId         string
		reserveFunc      func(ticketId string, data requests.Payload) error
		body             string
	}{
		"success: complete ticket purchase": {
			expectStatusCode: http.StatusOK,
			ticketId:         "id001",
			reserveFunc:      mockUpdateReserveTicketSuccess,
			body:             `{"userId": "teste123"}`,
		},
		// "failed: empty ticket Id": {
		// 	expectStatusCode: http.StatusBadRequest,
		// 	ticketId:         "",
		// },
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			reserveTicketFunc = test.reserveFunc

			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", fmt.Sprintf("/reserve/%s", test.ticketId), strings.NewReader(test.body))

			PostReserveTicket(w, r)
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
