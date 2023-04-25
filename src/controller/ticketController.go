package controllers

import (
	"a/src/models"
	"a/src/requests"
	"a/src/service"
	"context"
	"time"
)

type TicketController struct {
	ticketService service.TicketServiceInterface
}

func New(ticketService service.TicketServiceInterface) TicketController {
	return TicketController{ticketService: ticketService}
}

func (controller *TicketController) GetById(ticketId string) (*models.Ticket, error) {

	ctx := context.Background()

	requestCtx, cancel := context.WithTimeout(ctx, 10*time.Second)

	defer cancel()

	ticketModel, err := controller.ticketService.GetById(ticketId, requestCtx)

	return ticketModel, err
}

func (controller *TicketController) ReserveForPay(userId string, data requests.Payload) error {
	ctx := context.Background()

	requestCtx, cancel := context.WithTimeout(ctx, 10*time.Second)

	defer cancel()

	if err := controller.ticketService.ReserveForPay(userId, data, requestCtx); err != nil {
		return err
	}

	return nil
}
