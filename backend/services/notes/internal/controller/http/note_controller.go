package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Ixorlive/tw_vk/backend/services/notes/internal/entity"
	"github.com/Ixorlive/tw_vk/backend/services/notes/internal/usecase"
	"github.com/gin-gonic/gin"
)

type NoteController struct {
	Service usecase.NoteService
}

func NewNoteController(service usecase.NoteService) *NoteController {
	return &NoteController{Service: service}
}

func (nc *NoteController) RegisterRoutes(router *gin.Engine) {
	router.POST("/notes", nc.createNoteHandler)
	router.PUT("/notes/:id", nc.updateNoteHandler)
	router.DELETE("/notes/:id", nc.deleteNoteHandler)
	router.GET("/notes/:id", nc.getNoteByIDHandler)
	router.GET("/users/:userID/notes", nc.getNotesByUserIDHandler)
	router.GET("/notes", nc.getNotesByDateRangeHandler)
}

func (nc *NoteController) createNoteHandler(c *gin.Context) {
	var note entity.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := nc.Service.CreateNote(c, note.UserID, note.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (nc *NoteController) updateNoteHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var note entity.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := nc.Service.UpdateNote(c, id, note.UserID, note.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (nc *NoteController) deleteNoteHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID, _ := strconv.Atoi(c.Query("userID"))

	if err := nc.Service.DeleteNote(c, id, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (nc *NoteController) getNoteByIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID, _ := strconv.Atoi(c.Query("userID"))

	result, err := nc.Service.GetNoteByID(c, id, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (nc *NoteController) getNotesByUserIDHandler(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("userID"))

	results, err := nc.Service.GetNotesByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}

func (nc *NoteController) getNotesByDateRangeHandler(c *gin.Context) {
	start, err := time.Parse(time.RFC3339, c.Query("start"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date"})
		return
	}
	end, err := time.Parse(time.RFC3339, c.Query("end"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date"})
		return
	}
	userID, _ := strconv.Atoi(c.Query("userID"))

	results, err := nc.Service.GetNotesByDateRange(c, userID, start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}
