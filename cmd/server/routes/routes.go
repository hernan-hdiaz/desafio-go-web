package routes

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/server/handler"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

// Requerimiento 4:
// Una vez desarrollado el servicio y el repositorio desarrollar el archivo routes.go
// en en el package que corresponda. Los endpoints deberán ser los siguientes:

// GET - "/ticket/getByCountry/:dest"
// GET - “/ticket/getAverage/:dest”

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db []domain.Ticket
}

func NewRouter(r *gin.Engine, db []domain.Ticket) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.buildTicketRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1/ticket")
}

func (r *router) buildTicketRoutes() {
	repo := tickets.NewRepository(r.db)
	service := tickets.NewService(repo)
	ticketHandler := handler.NewService(service)

	r.rg.GET("/getByCountry/:dest", ticketHandler.GetTicketsByCountry())
	r.rg.GET("/getAverage/:dest", ticketHandler.AverageDestination())
}
