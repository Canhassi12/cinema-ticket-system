package service

import (
	"a/src/database"
	"a/src/models"
	"a/src/repositories"
	"a/src/requests"
	"fmt"

	"github.com/scylladb/gocqlx/v2"
)

type TicketServiceInterface interface {
	GetById(ticketId string) (models.Ticket, error)
	ReserveForPay(ticketId string, data requests.Payload) error
	ReleaseTicket(session *gocqlx.Session) error
	GetAllTickets(movieId string) ([]models.Ticket, error)
	PayTicket(ticketId string, userId string) error
	CheckItsSameUser(session *gocqlx.Session, ticketId string, userId string) error
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

func (ticketService *TicketService) ReserveForPay(ticketId string, data requests.Payload) error {
	session := ticketService.db.Conn()

	err := ticketService.ticketRepository.Update(session, ticketId, data)

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

func (ticketService *TicketService) PayTicket(ticketId string, userId string) error {
	session := ticketService.db.Conn()

	if err := ticketService.CheckItsSameUser(session, ticketId, userId); err != nil {
		return err
	}

	// payment ...

	if err := ticketService.ticketRepository.FinishedStatus(session, ticketId); err != nil {
		return err
	}

	return nil
}

func (ticketService *TicketService) CheckItsSameUser(session *gocqlx.Session, ticketId string, userId string) error {
	ticketModel, err := ticketService.ticketRepository.GetById(session, ticketId)

	if err != nil {
		return err
	}

	if userId != ticketModel.UserId {
		return fmt.Errorf("sorry, this ticket belongs to someone else")
	}

	return nil
}
