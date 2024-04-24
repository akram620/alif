package handler

import (
	"github.com/akram620/alif/internal/service"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	engine  *gin.Engine
	service service.Events
}

func NewHandler(chatService service.Events) *Handler {
	return &Handler{
		gin.New(),
		chatService,
	}
}

func (s *Handler) Run(port string) {
	s.engine.Use(logger.SetLogger())
	s.engine.Use(gin.Recovery())

	s.engine.GET("/health", s.Health)

	general := s.engine.Group("/api/v1")

	events := general.Group("/events")
	{
		events.POST("", s.CreateEvent)
	}

	s.engine.Run(port)
}
