package api

import (
	"github.com/akram620/alif/internal/service"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine  *gin.Engine
	service service.EventsService
}

func NewServer(chatService service.EventsService) *Server {
	return &Server{
		gin.New(),
		chatService,
	}
}

func (s *Server) Run(endpoint string) {
	s.engine.Use(logger.SetLogger())
	s.engine.Use(gin.Recovery())

	s.engine.GET("/health", s.Health)

	s.engine.Run(endpoint)
}
