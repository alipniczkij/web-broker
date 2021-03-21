package handler

import (
	"github.com/alipniczkij/web-broker/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	queue := router.Group("/queue")
	{
		queue.GET("/", h.getValue)
		queue.PUT("/", h.putValue)
	}
	return router
}
