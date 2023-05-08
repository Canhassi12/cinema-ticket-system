package service

import (
	"a/src/database"
	"a/src/models"
	"a/src/repositories"
	"a/src/requests"

	"github.com/scylladb/gocqlx/v2"
)

type TicketServiceInterface interface {
	GetById(ticketId string) (models.Ticket, error)
	ReserveForPay(userId string, data requests.Payload) error
	ReleaseTicket(session *gocqlx.Session) error
	GetAllTickets(movieId string) ([]models.Ticket, error)
}

type TicketService struct {
	ticketRepository repositories.TicketRepositoryInterface
	db               database.ScyllaConnInterface
}

func New(ticketRepository repositories.TicketRepositoryInterface, db database.ScyllaConnInterface) TicketService {
	return TicketService{ticketRepository: ticketRepository, db: db}
}

func (ticketService *TicketService) GetById(ticketId string) (models.Ticket, error) {
	session := ticketService.db.Conn()

	defer session.Close()

	ticketModel, err := ticketService.ticketRepository.GetById(session, ticketId)

	return ticketModel, err
}

func (ticketService *TicketService) ReserveForPay(userId string, data requests.Payload) error {
	session := ticketService.db.Conn()

	err := ticketService.ticketRepository.Update(session, userId, data)

	if err != nil {
		return err
	}

	return nil
}

func (ticketService *TicketService) ReleaseTicket(session *gocqlx.Session) error {
	if err := ticketService.ticketRepository.FreeStatus(session); err != nil {
		return err
	}

	return nil
}

func (ticketService *TicketService) GetAllTickets(movieId string) ([]models.Ticket, error) {
	session := ticketService.db.Conn()

	ticketModels, err := ticketService.ticketRepository.GetAll(session, movieId)

	if err != nil {
		return ticketModels, err
	}

	return ticketModels, nil
}
