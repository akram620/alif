package api

import (
	"github.com/akram620/alif/internal/service"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine  *gin.Engine
	service service.Events
}

func NewServer(chatService service.Events) *Server {
	return &Server{
		gin.New(),
		chatService,
	}
}

func (s *Server) Run(endpoint string) {
	s.engine.Use(logger.SetLogger())
	s.engine.Use(gin.Recovery())

	s.engine.GET("/health", s.Health)

	general := s.engine.Group("/api/v1")

	events := general.Group("/events")
	{
		events.POST("", s.CreateEvent)
	}

	s.engine.Run(endpoint)
}
