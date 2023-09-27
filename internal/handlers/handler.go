package handlers

import (
	"TestProjectEthereum/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{Service: service}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/withdraw", h.Withdraw)

	return router
}
