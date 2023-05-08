package repositories

import (
	"a/src/models"
	"a/src/requests"
	"fmt"
	"time"

	"github.com/scylladb/gocqlx/v2"
)

type TicketRepositoryInterface interface {
	GetById(session *gocqlx.Session, ticketId string) (models.Ticket, error)
	Update(session *gocqlx.Session, ticketId string, data requests.Payload) error
	GetAll(session *gocqlx.Session, movieId string) ([]models.Ticket, error)
	FreeStatus(session *gocqlx.Session) error
	FinishedStatus(session *gocqlx.Session, ticketId string) error
}

type TicketRepository struct {
}

func (ticketRepository *TicketRepository) GetById(session *gocqlx.Session, ticketId string) (models.Ticket, error) {

	var ticketModel models.Ticket

	q := session.Query("SELECT ticket_id, price, seat, movie, status, user_id FROM go.tickets WHERE ticket_id=:ticket_id", []string{":ticket_id"}).
		BindMap(map[string]interface{}{":ticket_id": ticketId})

	if err := q.GetRelease(&ticketModel); err != nil {
		return ticketModel, err
	}

	return ticketModel, nil
}

func (ticketRepository *TicketRepository) Update(session *gocqlx.Session, ticketId string, data requests.Payload) error {
	ticketModel := models.Ticket{
		Movie:  data.Movie,
		Seat:   12,
		Price:  models.FromPrice(data.Price),
		Status: models.Pending,
		UserId: data.UserId,
	}

	q := session.Query(
		`UPDATE go.tickets SET 
		   	status = :status,
		    user_id = :user_id,
			timestamp = :timestamp
			WHERE ticket_id = :ticket_id
			IF status ='free'`,
		[]string{":status", ":user_id", ":timestamp", ":ticket_id"}).
		BindMap(map[string]interface{}{
			":status":    ticketModel.Status,
			":user_id":   ticketModel.UserId,
			":timestamp": time.Now(),
			":seat":      ticketModel.Seat,
			":ticket_id": ticketId,
		})

	res, err := q.ExecCASRelease()

	if err != nil {
		return fmt.Errorf("error in exec update query: %w", err)
	}

	if !res {
		return fmt.Errorf("this ticket already be reserved")
	}

	return nil
}

func (ticketRepository *TicketRepository) GetAll(session *gocqlx.Session, movieId string) ([]models.Ticket, error) {
	ticketModel := []models.Ticket{}

	q := session.Query("SELECT * FROM go.tickets WHERE movie = :movie", []string{":movie"}).
		BindMap(map[string]interface{}{":movie": movieId})

	if err := q.SelectRelease(&ticketModel); err != nil {
		return ticketModel, fmt.Errorf("error in exec get all tickets query: %w", err)
	}

	return ticketModel, nil
}

func (ticketRepository *TicketRepository) GetPending(session *gocqlx.Session) ([]models.Ticket, error) {
	ticketModel := []models.Ticket{}

	q := session.Query("SELECT * FROM go.tickets WHERE status = 'pending'", []string{})

	if err := q.SelectRelease(&ticketModel); err != nil {
		return ticketModel, fmt.Errorf("error in exec get pending tickets query: %w", err)
	}

	return ticketModel, nil
}

func (ticketRepository *TicketRepository) FreeStatus(session *gocqlx.Session) error {
	ticketModels, err := ticketRepository.GetPending(session)

	if err != nil {
		return err
	}

	for _, v := range ticketModels {

		diff := time.Now().After(v.Timestamp.Add(time.Minute))

		if !diff {
			return nil
		}

		q := session.Query(
			`UPDATE go.tickets SET 
				status = 'free',
				user_id = '',
				timestamp = ''
				WHERE ticket_id = :ticket_id`,
			[]string{":ticket_id"}).
			BindMap(map[string]interface{}{
				":ticket_id": v.Ticket_id,
			})

		if err := q.ExecRelease(); err != nil {
			return fmt.Errorf("error in exec free status query: %w, ticket_id: %s", err, v.Ticket_id)
		}

		fmt.Printf("the ticket %s, it was released", v.Ticket_id)
	}

	return nil
}

func (ticketRepository *TicketRepository) FinishedStatus(session *gocqlx.Session, ticketId string) error {
	q := session.Query(
		`UPDATE go.tickets SET 
			status = 'finished',
			timestamp = :timestamp
			WHERE ticket_id = :ticket_id`,
		[]string{":timestamp", ":ticket_id"}).
		BindMap(map[string]interface{}{
			":timestamp": time.Now(),
			":ticket_id": ticketId,
		})

	if err := q.ExecRelease(); err != nil {
		return fmt.Errorf("error in exec finished status query: %w, ticket_id: %s", err, ticketId)
	}

	return nil
}
