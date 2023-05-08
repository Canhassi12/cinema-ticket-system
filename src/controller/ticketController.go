package controllers

import (
	"a/src/models"
	"a/src/requests"
	"a/src/service"
)

type TicketController struct {
	ticketService service.TicketServiceInterface
}

func New(ticketService service.TicketServiceInterface) TicketController {
	return TicketController{ticketService: ticketService}
}

func (controller *TicketController) GetById(ticketId string) (models.Ticket, error) {
	ticketModel, err := controller.ticketService.GetById(ticketId)

	return ticketModel, err
}

func (controller *TicketController) ReserveForPay(userId string, data requests.Payload) error {
	if err := controller.ticketService.ReserveForPay(userId, data); err != nil {
		return err
	}

	return nil
}

func (controller *TicketController) GetAll(movieId string) ([]models.Ticket, error) {
	ticketModels, err := controller.ticketService.GetAllTickets(movieId)

	if err != nil {
		return ticketModels, nil
	}

	return ticketModels, nil
}
