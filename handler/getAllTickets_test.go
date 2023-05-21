package handler

import (
	"a/src/models"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockGetAllTicketsSuccessfunc(movieId string) ([]models.Ticket, error) {
	return []models.Ticket{}, nil
}

func TestGetAllTickets(t *testing.T) {
	tests := map[string]struct {
		expectStatusCode int
		movie            string
		getFunc          func(movieId string) ([]models.Ticket, error)
	}{
		"success: got all tickets by movie": {
			expectStatusCode: http.StatusOK,
			movie:            "opa",
			getFunc:          mockGetAllTicketsSuccessfunc,
		},
		// "failed: empty ticket Id": {
		// 	expectStatusCode: http.StatusBadRequest,
		// 	ticketId:         "",
		// },
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			getAllTicketsFunc = test.getFunc

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", fmt.Sprintf("/tickets/%s", test.movie), nil)

			GetAllTickets(w, r)
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
