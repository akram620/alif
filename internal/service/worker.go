package service

import (
	"context"
	"github.com/akram620/alif/internal/models"
	"github.com/akram620/alif/internal/repository"
	"github.com/akram620/alif/pkg/logger"
	"time"
)

type Worker struct {
	eventsRepository repository.Events
	notifier         Notifier
}

func NewWorkerService(eventsRepository repository.Events, notifier Notifier) *Worker {
	return &Worker{eventsRepository, notifier}
}

func (w *Worker) RunJobs(ctx context.Context, duration time.Duration) {
	ticker := time.NewTicker(duration)
	for {
		select {
		case <-ctx.Done():
			logger.Info("🛑 Stopping Worker...")
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
		err := w.notifier.SendNotification(&event)
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
