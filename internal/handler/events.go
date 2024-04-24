package handler

import (
	"github.com/akram620/alif/internal/models"
	"github.com/akram620/alif/pkg/errors"
	"github.com/akram620/alif/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Handler) CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.BindJSON(&event); err != nil {
		logger.Error("CreateEvent: %v", err)
		c.JSON(http.StatusBadRequest, errors.NewErrorResponse(
			errors.ErrBadRequestParseBody.WithMessage(err.Error())),
		)
		return
	}

	if err := s.service.CreateEvent(&event); err != nil {
		logger.Error("CreateEvent: %v", err)
		c.JSON(err.HttpStatus, errors.NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, errors.NewSuccessResponse())
}
