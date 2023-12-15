package handler

import (
	"github.com/bootcamp-go/desafio-go-web/internal/ticket"
	"github.com/bootcamp-go/desafio-go-web/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TicketRouter struct {
	ticketGroup *gin.RouterGroup
	service     ticket.Service
}

func (s *TicketRouter) Routes() {
	//Get all tickets
	s.ticketGroup.GET("/", s.GetAll())
	//Get tickets by destination
	s.ticketGroup.GET("/getByCountry/:dest", s.GetTicketsByCountry())
	//Get average by destination
	s.ticketGroup.GET("/getAverage/:dest", s.AverageDestination())
	//Get total tickets by destination
	s.ticketGroup.GET("/getTotal/:dest", s.GetTotalTickets())
}

func NewTicketRouter(g *gin.RouterGroup) TicketRouter {
	slice, err := pkg.LoadTicketsFromFile("../tickets.csv")
	if err != nil {
		panic(err)
	}
	repo := ticket.NewRepository(slice)

	serv := ticket.NewTicketService(repo)

	return TicketRouter{g, serv}
}

func (s *TicketRouter) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		tickets, err := s.service.GetAll(c)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}
		c.JSON(200, tickets)
	}
}

func (s *TicketRouter) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := s.service.GetTicketByDestination(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tickets)
	}
}

func (s *TicketRouter) GetTotalTickets() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		total, err := s.service.GetTotalTickets(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, total)
	}
}

func (s *TicketRouter) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := s.service.AverageDestination(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, avg)
	}
}
