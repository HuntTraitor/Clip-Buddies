package system

import "github.com/gin-gonic/gin"

type Handler struct {
	service *Service
}

func NewHandler() *Handler {
	service := NewService()

	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterRoutes(group *gin.RouterGroup) {
	group.GET("", h.healthCheckHandler)
}

func (h *Handler) healthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": h.service.HealthStatus(),
	})
}
