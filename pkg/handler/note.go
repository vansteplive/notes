package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vansteplive/notes-app-golang/models"
)

func (h *Handler) CreateNote(c *gin.Context) {
	var input models.NoteItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Note.CreateNote(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetAllNotes(c *gin.Context) {
	list, err := h.services.Note.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"notes": list,
	})

}
