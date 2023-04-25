package service

import (
	"a/src/database"
	"a/src/models"
	"a/src/repositories"
	"a/src/requests"
	"context"
	"fmt"

	"github.com/scylladb/scylla-go-driver"
)

type TicketServiceInterface interface {
	GetById(ticketId string, requestCtx context.Context) (*models.Ticket, error)
	ReserveForPay(userId string, data requests.Payload, requestCtx context.Context) error
}

type TicketService struct {
	ticketRepository repositories.TicketRepositoryInterface
	db               database.ScyllaConnInterface
}

func New(ticketRepository repositories.TicketRepositoryInterface, db database.ScyllaConnInterface) TicketService {
	return TicketService{ticketRepository: ticketRepository, db: db}
}

func (ticketService *TicketService) GetById(ticketId string, requestCtx context.Context) (*models.Ticket, error) {
	session := ticketService.db.Conn()

	defer session.Close()

	ticketModel, err := ticketService.ticketRepository.GetById(requestCtx, session, ticketId)

	return ticketModel, err
}

func (ticketService *TicketService) ReserveForPay(userId string, data requests.Payload, requestCtx context.Context) error {
	session := ticketService.db.Conn()

	if err := ticketService.CheckPendences(requestCtx, session, userId, data.TicketId); err != nil {
		return err
	}

	err := ticketService.ticketRepository.Update(requestCtx, session, userId, data)

	if err != nil {
		return err
	}

	return nil
}

func (ticketService *TicketService) CheckPendences(requestCtx context.Context, session *scylla.Session, userId string, ticketId string) error {
	ticket, err := ticketService.ticketRepository.GetById(requestCtx, session, ticketId)

	if err != nil {
		return err
	}

	itsDifferentUser := ticket.UserId != userId

	if !itsDifferentUser {
		return fmt.Errorf("Sorry, another user is buying")
	}

	itsPending := ticket.Status == "pending"

	if itsPending {
		return fmt.Errorf("This ticket already be pending")
	}

	return nil
}
