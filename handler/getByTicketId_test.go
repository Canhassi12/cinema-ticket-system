package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"ticket-system/src/models"
)

func mockGetTicketByIdSuccess(ticketId string) (models.Ticket, error) {
	return models.Ticket{}, nil
}

func TestGetTicketById(t *testing.T) {
	tests := map[string]struct {
		expectStatusCode int
		ticketId         string
		getFunc          func(ticketId string) (models.Ticket, error)
	}{
		"success: got ticket Id": {
			expectStatusCode: http.StatusOK,
			ticketId:         "id001",
			getFunc:          mockGetTicketByIdSuccess,
		},
		// "failed: empty ticket Id": {
		// 	expectStatusCode: http.StatusBadRequest,
		// 	ticketId:         "",
		// },
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			getByIdFunc = test.getFunc

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", fmt.Sprintf("/ticket/%s", test.ticketId), nil)

			GetTicketById(w, r)
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
