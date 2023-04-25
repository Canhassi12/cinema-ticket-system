package repositories

import (
	"a/src/models"
	"a/src/requests"
	"context"
	"fmt"
	"time"

	"github.com/scylladb/scylla-go-driver"
)

const Table = "tickets"

type TicketRepositoryInterface interface {
	GetById(requestCtx context.Context, session *scylla.Session, ticketId string) (*models.Ticket, error)
	Update(requestCtx context.Context, session *scylla.Session, userId string, data requests.Payload) error
}

type TicketRepository struct {
}

func (*TicketRepository) GetById(requestCtx context.Context, session *scylla.Session, ticketId string) (*models.Ticket, error) {
	q, err := session.Prepare(requestCtx, "SELECT id, price, seat, movie, status, user_id FROM tickets WHERE id="+ticketId)

	if err != nil {
		return nil, err
	}

	res, err := q.Exec(requestCtx)

	if err != nil {
		return nil, err
	}

	// id, err := res.Rows[0][1].AsUUID()

	// if err != nil {
	// 	return nil, err
	// }

	price, err := res.Rows[0][1].AsInt32()

	if err != nil {
		return nil, err
	}

	seat, err := res.Rows[0][2].AsInt32()

	if err != nil {
		return nil, err
	}

	movie, err := res.Rows[0][3].AsText()

	if err != nil {
		return nil, err
	}

	status, err := res.Rows[0][4].AsText()

	if err != nil {
		return nil, err
	}

	userId, err := res.Rows[0][5].AsText()

	if err != nil {
		return nil, err
	}

	ticketModel := models.Ticket{Price: int(price), Movie: movie, Seat: int(seat), Status: status, UserId: userId}

	if err != nil {
		return nil, err
	}

	return &ticketModel, nil
}

func (ticketRepository *TicketRepository) Update(requestCtx context.Context, session *scylla.Session, userId string, data requests.Payload) error {

	ticketModel := models.Ticket{
		Movie:  data.Movie,
		Seat:   2,
		Price:  models.FromPrice(data.Price),
		Status: models.Pending,
		UserId: userId,
	}

	q, err := session.Prepare(requestCtx, fmt.Sprintf(
		"UPDATE %s SET movie = %s seat = %d price = %d status = %s user_id = %s",
		Table,
		ticketModel.Movie,
		ticketModel.Seat,
		ticketModel.Price,
		ticketModel.Status,
		ticketModel.UserId),
	)

	if err != nil {
		return err
	}

	res, err := q.Exec(requestCtx)

	if err != nil {
		return err
	}

	fmt.Printf("%v", res)

	go ticketRepository.FreeStatus(requestCtx, session, ticketModel)

	return nil
}

func (*TicketRepository) FreeStatus(
	requestCtx context.Context,
	session *scylla.Session,
	ticketModel models.Ticket,
) {
	time.Sleep(time.Minute * 10)

	_, err := session.Prepare(requestCtx, fmt.Sprintf(
		"UPDATE %s SET movie = %s seat = %d price = %d status = %s user_id = null",
		Table,
		ticketModel.Movie,
		ticketModel.Seat,
		ticketModel.Price,
		models.Free),
	)

	if err != nil {
		value, _ := fmt.Printf("problems to free this %d seat, error: %s", ticketModel.Seat, err.Error())

		panic(value)
	}
}
