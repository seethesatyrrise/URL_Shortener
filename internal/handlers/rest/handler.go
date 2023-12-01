package rest

import (
	"URL_Shortener/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Route() *gin.Engine {
	router := gin.Default()
	{
		router.GET(":token", h.GetFullLink)
		router.POST("/api/generate", h.GetToken)
	}
	return router
}
