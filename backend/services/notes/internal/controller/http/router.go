package http

import (
	"strings"
	"time"

	"github.com/Ixorlive/tw_vk/backend/services/notes/internal/usecase"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Ixorlive/tw_vk/backend/services/notes/docs"
)

func NewRouter(noteService usecase.NoteService, cors_allow_origins string) *gin.Engine {
	router := gin.Default()
	// swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	router.GET("/swagger/*any", swaggerHandler)

	config := cors.Config{
		AllowOrigins:     strings.Split(cors_allow_origins, ","),
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(config))
	// router.Use(cors.Default())
	noteController := NewNoteController(noteService)
	noteController.RegisterRoutes(router)
	return router
}
