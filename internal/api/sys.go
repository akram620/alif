package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health		 godoc
// @Summary      Return health status
// @Tags		 System
// @Produce      json
// @Success		 200 {object} models.StatusResponse
// @Router       /market/health [get]
func (s *Server) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
