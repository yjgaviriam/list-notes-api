package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"list-notes/app/models"
	"list-notes/app/repositories"
	"net/http"
)

type NoteController struct {
	DB *sql.DB
}

func NewNoteController(db *sql.DB) NoteControllerInterface {
	return &NoteController{DB: db}
}

// GetNote implements NoteControllerInterface
func (n *NoteController) GetNote(g *gin.Context) {
	db := n.DB
	repositoryNote := repositories.NewNoteRepository(db)
	getNote := repositoryNote.SelectNote()
	if getNote != nil {
		g.JSON(http.StatusOK, gin.H{"status": "success", "data": getNote, "msg": "get note successfully"})
	} else {
		g.JSON(http.StatusOK, gin.H{"status": "success", "data": nil, "msg": "get note successfully"})
	}
}

// InsertNote implements NoteControllerInterface
func (n *NoteController) InsertNote(g *gin.Context) {
	db := n.DB
	var post models.PostNote
	if err := g.ShouldBindJSON(&post); err == nil {
		repositoryNote := repositories.NewNoteRepository(db)
		insertNote := repositoryNote.InsertNote(post)
		if insertNote {
			g.JSON(http.StatusOK, gin.H{"status": "success", "msg": "insert note successfully"})
		} else {
			g.JSON(http.StatusBadRequest, gin.H{"status": "failed", "msg": "insert note failed"})
		}
	} else {
		g.JSON(http.StatusBadRequest, gin.H{"status": "success", "msg": err})
	}
}
