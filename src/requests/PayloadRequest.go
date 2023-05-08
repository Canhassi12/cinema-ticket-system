package requests

type Payload struct {
	Price    string `json:"price"`
	Movie    string `json:"movie"`
	Seat     int    `json:"seat"`
	TicketId string `json:"ticketId"`
	UserId   string `json:"userId"`
}
