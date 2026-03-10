package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hunttraitor/clip-buddies/internal/system"
)

type SystemHandler struct {
	SystemService *system.Service
}

func NewSystemHandler(systemService *system.Service) *SystemHandler {
	return &SystemHandler{
		SystemService: systemService,
	}
}

func (h *SystemHandler) RegisterRoutes(group *gin.RouterGroup) {
	group.GET("", h.HealthcheckHandler)
}

func (h *SystemHandler) HealthcheckHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "up"})
}
