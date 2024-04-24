package service

import (
	"github.com/akram620/alif/internal/models"
	"github.com/akram620/alif/pkg/logger"
)

type Notification interface {
	SendNotification(event *models.Event) error
}

type NotificationService struct{}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (n *NotificationService) SendNotification(event *models.Event) error {
	logger.Infof("📨 Sending notification for event: %v", event)
	return nil
}
