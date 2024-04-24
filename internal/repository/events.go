package repository

import (
	"context"
	"github.com/akram620/alif/internal/errors"
	"github.com/akram620/alif/internal/logger"
	"github.com/akram620/alif/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Events interface {
	CreateEvent(e *models.Event) *errors.ExportableError
}

type EventsRepository struct {
	pool *pgxpool.Pool
}

func NewEventsRepository(pool *pgxpool.Pool) *EventsRepository {
	return &EventsRepository{pool}
}

func (r *EventsRepository) CreateEvent(e *models.Event) *errors.ExportableError {
	query := `
		INSERT INTO events (order_type, session_id, card, event_date, website_url)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.pool.Exec(context.Background(), query, e.OrderType, e.SessionID, e.Card, e.EventDate, e.WebsiteURL)
	if err != nil {
		logger.Error("EventsRepository.CreateEvent: %v", err)
		return &errors.ErrInternalServerErrorDatabaseFailed
	}

	return nil
}
