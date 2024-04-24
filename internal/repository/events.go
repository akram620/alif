package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Events interface {
}

type EventsRepository struct {
	pool *pgxpool.Pool
}

func NewEventsRepository(pool *pgxpool.Pool) *EventsRepository {
	return &EventsRepository{pool}
}
