package auth

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(db *sql.DB) *Handler {
	repo := NewRepository(db)
	service := NewService(repo)

	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterRoutes(group *gin.RouterGroup) {
	group.GET("/auth/ping", h.pingHandler)
}

func (h *Handler) pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "auth ok",
	})
}
