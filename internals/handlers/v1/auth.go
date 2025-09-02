package v1

import (
	"gorestapi/internals/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

func (h *AuthHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/login", h.Login)
	router.POST("/register", h.Register)
}
