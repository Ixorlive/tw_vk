package http

import (
	"github.com/Ixorlive/tw_vk/backend/services/auth/internal/usecase"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Ixorlive/tw_vk/backend/services/auth/docs"
)

func NewRouter(authService usecase.AuthService) *gin.Engine {
	router := gin.Default()
	// swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	router.GET("/swagger/*any", swaggerHandler)

	controller := NewAuthController(authService)

	router.POST("/login", controller.AuthByLogin)
	router.POST("/token", controller.AuthByToken)
	router.POST("/register", controller.Register)

	return router
}
