package service

import (
	"github.com/akram620/alif/internal/repository"
)

type Events interface {
}

type EventsService struct {
	eventsRepository repository.Events
}

func NewEventsService(chatRepository repository.Events) *EventsService {
	return &EventsService{chatRepository}
}
