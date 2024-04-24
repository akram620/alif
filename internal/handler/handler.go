package handler

import (
	"github.com/akram620/alif/internal/service"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Events
}

func NewHandler(service service.Events) *Handler {
	return &Handler{
		service,
	}
}

func (s *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(logger.SetLogger())
	router.Use(gin.Recovery())

	router.GET("/health", s.Health)

	general := router.Group("/api/v1")

	events := general.Group("/events")
	{
		events.POST("", s.CreateEvent)
	}

	return router
}
