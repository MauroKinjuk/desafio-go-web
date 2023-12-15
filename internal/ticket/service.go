package ticket

import (
	"context"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTotalTickets(ctx context.Context, destination string) (float64, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type TicketService struct {
	TicketService TicketRepository
}

func (s *TicketService) Ping() gin.HandlerFunc {
	return s.Ping()
}

func NewTicketService(repo TicketRepository) Service {
	return &TicketService{repo}
}

func (s *TicketService) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	tickets, err := s.TicketService.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (s *TicketService) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {
	return s.TicketService.GetTicketByDestination(ctx, destination)
}

func (s *TicketService) GetTotalTickets(ctx context.Context, destination string) (float64, error) {
	return s.TicketService.GetTotalTickets(ctx, destination)
}

func (s *TicketService) AverageDestination(ctx context.Context, destination string) (float64, error) {
	return s.TicketService.AverageDestination(ctx, destination)
}
