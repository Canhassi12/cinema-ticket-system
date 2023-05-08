package models

import "time"

type Ticket struct {
	Ticket_id string
	Price     int
	Movie     string `json:"movie"`
	Seat      int    `json:"seat"`
	Status    string
	UserId    string
	Timestamp time.Time
}

type TicketInterface interface {
}

var (
	HalfPrice = "half"
	FullPrice = "full"
)

func FromPrice(p string) int {
	switch p {
	case HalfPrice:
		return 15
	case FullPrice:
		return 30
	}

	return 30
}

var (
	Free     = "free"
	Pending  = "pending"
	Finished = "finished"
)

func FromStatus(s string) string {
	switch s {
	case Free:
		return "free"
	case Pending:
		return "pending"
	case Finished:
		return "finished"
	}

	return "free"
}
