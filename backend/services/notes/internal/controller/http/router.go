package http

import (
	"github.com/Ixorlive/tw_vk/backend/services/notes/internal/usecase"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Ixorlive/tw_vk/backend/services/notes/docs"
)

func NewRouter(noteService usecase.NoteService) *gin.Engine {
	router := gin.Default()
	// swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	router.GET("/swagger/*any", swaggerHandler)

	noteController := NewNoteController(noteService)
	noteController.RegisterRoutes(router)
	return router
}
