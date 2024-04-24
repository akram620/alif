package service

import (
	"github.com/akram620/alif/internal/models"
	"github.com/akram620/alif/pkg/logger"
)

type Notifier interface {
	SendNotification(event *models.Event) error
}

type NotifierService struct{}

func NewNotifierService() *NotifierService {
	return &NotifierService{}
}

func (n *NotifierService) SendNotification(event *models.Event) error {
	logger.Infof("📨 Sending notification for event: %v", event)
	return nil
}
