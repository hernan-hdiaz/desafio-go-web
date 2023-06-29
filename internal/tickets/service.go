package tickets

import (
	"context"
)

// Requerimiento 3:
// Utilizando el archivo service_test.go como guía, desarrollar la lógica de negocio en service.go.
// Desarrollar los métodos correspondientes a la estructura Ticket.
// Uno de ellos debe devolver la cantidad de tickets de un destino.
// El  otro método debe devolver el promedio de personas que viajan a un país determinado en un dia:

type Service interface {
	AverageDestination(ctx context.Context, destination string) (int, error)
	GetTotalTickets(ctx context.Context, destination string) (int, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) AverageDestination(ctx context.Context, destination string) (int, error) {

	// tickets al destino
	ticketsDest, err := s.GetTotalTickets(ctx, destination)
	if err != nil {
		return 0, err
	}

	// todos los tickets
	allTickets, err := s.repo.GetAll(ctx)
	if err != nil {
		return 0, err
	}

	// porcentaje en comparación con el total de tickets
	avr := ticketsDest * 100 / len(allTickets)

	return avr, nil
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) (int, error) {
	ticketsDest, err := s.repo.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}
	return len(ticketsDest), nil
}
