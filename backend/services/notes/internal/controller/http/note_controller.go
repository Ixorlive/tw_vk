package http

import (
	"net/http"
	"strconv"
	"time"

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
	router.GET("/notes", nc.getAllNotesHandler)
	router.POST("/notes", nc.createNoteHandler)
	router.PUT("/notes/:id", nc.updateNoteHandler)
	router.DELETE("/notes/:id", nc.deleteNoteHandler)
	router.GET("/notes/:id", nc.getNoteByIDHandler)
	router.GET("/users/:userID/notes", nc.getNotesByUserIDHandler)
	router.GET("/notes/range", nc.getNotesByDateRangeHandler)
}

// @Summary Get all notes
// @Description Get all notes from db
// @Tags Notes
// @Produce json
// @Success 201 {object} entity.Note "List of all notes"
// @Failure 400 {object} map[string]string "Bad request if the JSON body cannot be parsed"
// @Failure 500 {object} map[string]string "Internal Server Error if there is a problem creating the note"
// @Router /notes [get]
func (nc *NoteController) getAllNotesHandler(c *gin.Context) {
	res, err := nc.Service.GetNotes(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

type CreateNoteParam struct {
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
}

// @Summary Create a note
// @Description Creates a new note with the given content and user ID.
// @Tags Notes
// @Accept json
// @Produce json
// @Param note body CreateNoteParam true "Note Content"
// @Success 201 {object} entity.Note "Note successfully created"
// @Failure 400 {object} map[string]string "Bad request if the JSON body cannot be parsed"
// @Failure 500 {object} map[string]string "Internal Server Error if there is a problem creating the note"
// @Router /notes [post]
func (nc *NoteController) createNoteHandler(c *gin.Context) {
	var note CreateNoteParam
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

// @Summary Update a note
// @Description Updates an existing note identified by its ID with new content.
// @Tags Notes
// @Accept json
// @Produce json
// @Param id path int true "Note ID"
// @Param note body CreateNoteParam true "Updated Note Content"
// @Success 200 {object} entity.Note "Note successfully updated"
// @Failure 400 {object} map[string]string "Bad request if the JSON body cannot be parsed"
// @Failure 500 {object} map[string]string "Internal Server Error if there is a problem updating the note"
// @Router /notes/{id} [put]
func (nc *NoteController) updateNoteHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var note CreateNoteParam
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

// @Summary Delete a note
// @Description Deletes a note identified by its ID.
// @Tags Notes
// @Produce json
// @Param id path int true "Note ID"
// @Success 200 "Note successfully deleted"
// @Failure 500 {object} map[string]string "Internal Server Error if there is a problem deleting the note"
// @Router /notes/{id} [delete]
func (nc *NoteController) deleteNoteHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := nc.Service.DeleteNote(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Get a note by ID
// @Description Retrieves a note by its ID for a specific user.
// @Tags Notes
// @Produce json
// @Param id path int true "Note ID"
// @Param userID query int true "User ID of the note owner"
// @Success 200 {object} entity.Note "Note found and returned"
// @Failure 500 {object} map[string]string "Internal Server Error if there is a problem retrieving the note"
// @Router /notes/{id} [get]
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

// @Summary Get notes by user ID
// @Description Retrieves all notes for a specific user.
// @Tags Notes
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {array} entity.Note "List of notes owned by the user"
// @Failure 500 {object} map[string]string "Internal Server Error if there is a problem retrieving notes"
// @Router /users/{userID}/notes [get]
func (nc *NoteController) getNotesByUserIDHandler(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("userID"))

	results, err := nc.Service.GetNotesByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}

// @Summary Get notes by date range
// @Description Retrieves all notes for a specific user within a specified date range.
// @Tags Notes
// @Produce json
// @Param userID query int true "User ID"
// @Param start query string true "Start date (RFC3339 format)"
// @Param end query string true "End date (RFC3339 format)"
// @Success 200 {array} entity.Note "List of notes within the date range"
// @Failure 400 {object} map[string]string "Bad request if date parameters are invalid"
// @Failure 500 {object} map[string]string "Internal Server Error if there is a problem retrieving notes"
// @Router /notes_range [get]
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
