package service

import (
	"github.com/akram620/alif/internal/models"
	"github.com/akram620/alif/internal/repository"
	"github.com/akram620/alif/pkg/errors"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.3 --name=Events

type Events interface {
	CreateEvent(e *models.Event) *errors.ExportableError
}

func (s *EventsService) CreateEvent(e *models.Event) *errors.ExportableError {
	return s.eventsRepository.CreateEvent(e)
}

type EventsService struct {
	eventsRepository repository.Events
}

func NewEventsService(chatRepository repository.Events) *EventsService {
	return &EventsService{chatRepository}
}
