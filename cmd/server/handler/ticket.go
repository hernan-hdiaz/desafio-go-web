package handler

import (
	"errors"
	"net/http"

	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

type Service struct {
	service tickets.Service
}

func NewService(s tickets.Service) *Service {
	return &Service{
		service: s,
	}
}

func (s *Service) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		totalTickets, err := s.service.GetTotalTickets(c, destination)
		if err != nil {
			switch {
			case errors.Is(err, tickets.ErrNotFound):
				c.String(http.StatusNotFound, err.Error())
			case errors.Is(err, tickets.ErrEmptyList):
				c.String(http.StatusInternalServerError, err.Error())
			}
			return
		}

		c.JSON(http.StatusOK, totalTickets)
	}
}

func (s *Service) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := s.service.AverageDestination(c, destination)
		if err != nil {
			switch {
			case errors.Is(err, tickets.ErrNotFound):
				c.String(http.StatusNotFound, err.Error())
			case errors.Is(err, tickets.ErrEmptyList):
				c.String(http.StatusInternalServerError, err.Error())
			}
			return
		}

		c.JSON(http.StatusOK, avg)
	}
}
