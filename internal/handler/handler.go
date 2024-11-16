package handler

import (
	"doodocs_rest_api/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	archive := router.Group("/api/archive")
	{
		archive.POST("/information", h.ArchiveInfo)
	}

	return router
}
