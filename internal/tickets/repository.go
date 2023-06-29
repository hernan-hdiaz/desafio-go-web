package tickets

import (
	"context"
	"errors"

	"github.com/bootcamp-go/desafio-go-web/internal/domain"
)

var (
	ErrEmptyList = errors.New("empty list of tickets")
	ErrNotFound  = errors.New("destination not found")
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
}

type repository struct {
	db []domain.Ticket
}

func NewRepository(db []domain.Ticket) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Ticket, error) {

	if len(r.db) == 0 {
		return []domain.Ticket{}, ErrEmptyList
	}

	return r.db, nil
}

func (r *repository) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {

	var ticketsDest []domain.Ticket

	if len(r.db) == 0 {
		return []domain.Ticket{}, ErrEmptyList
	}

	for _, t := range r.db {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	if len(ticketsDest) == 0 {
		return []domain.Ticket{}, ErrNotFound
	}

	return ticketsDest, nil
}
