package http

import (
	"github.com/Ixorlive/tw_vk/backend/services/auth/internal/controller/http/handlers"
	"github.com/Ixorlive/tw_vk/backend/services/auth/internal/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(authService usecase.AuthService) *gin.Engine {
	router := gin.Default()

	handler := handlers.NewAuthHandler(authService)

	router.POST("/login", handler.AuthByLogin)
	router.POST("/token", handler.AuthByToken)
	router.POST("/register", handler.Register)

	return router
}
