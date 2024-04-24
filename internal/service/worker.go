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
	for _, event := range *events {
		logger.Infof("📨 Sending notification for event: %+v", event)
	}
}
