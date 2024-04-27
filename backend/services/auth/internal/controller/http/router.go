package http

import (
	"github.com/Ixorlive/tw_vk/backend/services/auth/internal/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(authService usecase.AuthService) *gin.Engine {
	router := gin.Default()

	controller := NewAuthController(authService)

	router.POST("/login", controller.AuthByLogin)
	router.POST("/token", controller.AuthByToken)
	router.POST("/register", controller.Register)

	return router
}
