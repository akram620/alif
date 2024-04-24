package repository

import (
	"context"
	"github.com/akram620/alif/internal/models"
	"github.com/akram620/alif/pkg/errors"
	"github.com/akram620/alif/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Events interface {
	CreateEvent(e *models.Event) *errors.ExportableError
	GetEvents() (*[]models.Event, *errors.ExportableError)
	MarkEventsAsSent(ids []int64) *errors.ExportableError
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

	eventDate, err := time.Parse("2006-01-02 15:04:05.999999 -07:00", e.EventDate)
	if err != nil {
		logger.Error(err)
		return &errors.ErrInternalServerErrorDatabaseFailed
	}
	eventDate = eventDate.UTC()

	_, err = r.pool.Exec(context.Background(), query, e.OrderType, e.SessionID, e.Card, eventDate, e.WebsiteURL)
	if err != nil {
		return &errors.ErrInternalServerErrorDatabaseFailed
	}

	return nil
}

func (r *EventsRepository) GetEvents() (*[]models.Event, *errors.ExportableError) {
	query := `
		SELECT id, order_type, session_id, card, event_date, website_url
		FROM events
		WHERE deleted_at is null and sent = false and
			date_trunc('minute', event_date) = date_trunc('minute', CURRENT_TIMESTAMP);
	`
	rows, err := r.pool.Query(context.Background(), query)
	if err != nil {
		logger.Error(err)
		return nil, &errors.ErrInternalServerErrorDatabaseFailed
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var e models.Event
		var evDate time.Time
		err = rows.Scan(&e.ID, &e.OrderType, &e.SessionID, &e.Card, &evDate, &e.WebsiteURL)
		if err != nil {
			logger.Error(err)
			return nil, &errors.ErrInternalServerErrorDatabaseFailed
		}

		e.EventDate = evDate.Format(time.RFC3339)
		events = append(events, e)
	}

	return &events, nil
}

func (r *EventsRepository) MarkEventsAsSent(ids []int64) *errors.ExportableError {
	query := `
		UPDATE events
		SET sent = true
		WHERE id = ANY($1)
	`
	_, err := r.pool.Exec(context.Background(), query, ids)
	if err != nil {
		logger.Error(err)
		return &errors.ErrInternalServerErrorDatabaseFailed
	}

	return nil
}
