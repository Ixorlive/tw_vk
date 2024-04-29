package http

import (
	"fmt"
	"strings"
	"time"

	_ "github.com/Ixorlive/tw_vk/backend/services/auth/docs"
	"github.com/Ixorlive/tw_vk/backend/services/auth/internal/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(authService usecase.AuthService, cors_allow_origins string) *gin.Engine {
	router := gin.Default()
	// swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	router.GET("/swagger/*any", swaggerHandler)

	fmt.Println(cors_allow_origins)
	config := cors.Config{
		AllowOrigins:     strings.Split(cors_allow_origins, ","),
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(config))

	controller := NewAuthController(authService)

	router.POST("/login", controller.AuthByLogin)
	router.POST("/token", controller.AuthByToken)
	router.POST("/register", controller.Register)

	return router
}
