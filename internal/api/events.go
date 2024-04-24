package api

import (
	"github.com/akram620/alif/internal/errors"
	"github.com/akram620/alif/internal/logger"
	"github.com/akram620/alif/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.BindJSON(&event); err != nil {
		logger.Error("CreateEvent: %v", err)
		c.JSON(http.StatusBadRequest, models.NewErrorResponse(&errors.ErrBadRequestParseBody))
		return
	}

	if err := s.service.CreateEvent(&event); err != nil {
		logger.Error("CreateEvent: %v", err)
		c.JSON(err.HttpStatus, models.NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.NewSuccessResponse())
}
