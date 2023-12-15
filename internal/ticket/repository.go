package ticket

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/bootcamp-go/desafio-go-web/internal/domain"
)

type TicketRepository interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
	GetTotalTickets(ctx context.Context, destination string) (float64, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type repositoryTicket struct {
	ticketsDB []domain.Ticket
}

func NewRepository(db []domain.Ticket) TicketRepository {
	return &repositoryTicket{
		ticketsDB: db}
}

func (r *repositoryTicket) Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}

func (r *repositoryTicket) GetAll(ctx context.Context) ([]domain.Ticket, error) {

	if len(r.ticketsDB) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	return r.ticketsDB, nil
}

func (r *repositoryTicket) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {

	var ticketsDest []domain.Ticket

	if len(r.ticketsDB) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.ticketsDB {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	return ticketsDest, nil
}

func (r *repositoryTicket) GetTotalTickets(ctx context.Context, destination string) (float64, error) {

	ticketsDest := 0.0

	if len(r.ticketsDB) == 0 {
		return 0, fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.ticketsDB {
		if t.Country == destination {
			ticketsDest++
		}
	}

	return ticketsDest, nil
}

// Func for Average destination, return percentage of tickets per destination
func (r *repositoryTicket) AverageDestination(ctx context.Context, destination string) (float64, error) {

	ticketsDest := 0.0
	totalTickets := 0.0

	if len(r.ticketsDB) == 0 {
		return 0, fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.ticketsDB {
		if t.Country == destination {
			ticketsDest++
		}
		totalTickets++
	}

	return (ticketsDest / totalTickets) * 100, nil
}
