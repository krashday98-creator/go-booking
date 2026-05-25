package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthResponse ответ проверки состояния сервиса.
type HealthResponse struct {
	Status string `json:"status" example:"ok"`
}

// Health проверяет доступность API.
//
// @Summary      Health check
// @Tags         system
// @Produce      json
// @Success      200 {object} HealthResponse
// @Router       /health [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{Status: "ok"})
}
