package service

import (
	"context"
	"github.com/akram620/alif/internal/logger"
	"github.com/akram620/alif/internal/models"
	"github.com/akram620/alif/internal/repository"
	"time"
)

type Worker struct {
	eventsRepository repository.Events
}

func NewWorkerService(eventsRepository repository.Events) *Worker {
	return &Worker{eventsRepository}
}

func (w *Worker) RunJobs(ctx context.Context, duration time.Duration) {
	ticker := time.NewTicker(duration)
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			return
		case <-ticker.C:
			logger.Info("⏰ Ticks Received...")
			events, err := w.eventsRepository.GetEvents()
			if err != nil {
				logger.Error(err)
				continue
			}
			// отправляем уведомления
			w.sendNotifications(events)
		}

	}
}

func (w *Worker) sendNotifications(events *[]models.Event) {
	success := make([]int64, 0, len(*events))
	for _, event := range *events {
		err := w.sendNotification(&event)
		if err != nil {
			logger.Error(err)
			continue
		}

		success = append(success, event.ID)
	}

	if len(success) > 0 {
		err := w.eventsRepository.MarkEventsAsSent(success)
		if err != nil {
			logger.Error(err)
		}
	}
}

func (w *Worker) sendNotification(event *models.Event) error {
	logger.Infof("📨 Sending notification for event: %v", event)
	return nil
}
