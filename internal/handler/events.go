package handler

import (
	"github.com/akram620/alif/internal/models"
	errors2 "github.com/akram620/alif/pkg/errors"
	"github.com/akram620/alif/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Handler) CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.BindJSON(&event); err != nil {
		logger.Error("CreateEvent: %v", err)
		c.JSON(http.StatusBadRequest, errors2.NewErrorResponse(&errors2.ErrBadRequestParseBody))
		return
	}

	if err := s.service.CreateEvent(&event); err != nil {
		logger.Error("CreateEvent: %v", err)
		c.JSON(err.HttpStatus, errors2.NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, errors2.NewSuccessResponse())
}
