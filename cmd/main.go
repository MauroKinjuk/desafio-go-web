package main

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	ticketGroup := r.Group("/tickets")
	ticketRouter := handler.NewTicketRouter(ticketGroup)
	ticketRouter.Routes()

	if err := r.Run(); err != nil {
		panic(err)
	}

}
