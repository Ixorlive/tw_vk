package http

import (
	"github.com/Ixorlive/tw_vk/backend/services/notes/internal/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(noteService usecase.NoteService) *gin.Engine {
	router := gin.Default()
	noteController := NewNoteController(noteService)
	noteController.RegisterRoutes(router)
	return router
}
