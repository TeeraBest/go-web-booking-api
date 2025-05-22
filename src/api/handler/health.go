package handler

import (
	"net/http"

	"go-web-api/api/helper"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthCheck godoc
// @Summary Health Check
// @Description Health Check
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/health/ [get]
func (h *HealthHandler) Health(c *gin.Context) {
	go func() {
		// Simulate some background processing (e.g., logging, metrics, etc.)
		// Example: log.Println("Background task running...")
		// Respond immediately
		c.JSON(http.StatusOK, helper.GenerateBaseResponse("Working!", true, 0))
	}()

}
