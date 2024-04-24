package service

import (
	"github.com/akram620/alif/internal/models"
	"github.com/akram620/alif/internal/repository"
	"github.com/akram620/alif/pkg/errors"
)

type Events interface {
	CreateEvent(e *models.Event) *errors.ExportableError
}

type EventsService struct {
	eventsRepository repository.Events
}

func NewEventsService(chatRepository repository.Events) *EventsService {
	return &EventsService{chatRepository}
}

func (s *EventsService) CreateEvent(e *models.Event) *errors.ExportableError {
	return s.eventsRepository.CreateEvent(e)
}
