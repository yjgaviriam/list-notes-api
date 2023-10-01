package controllers

import "github.com/gin-gonic/gin"

type NoteControllerInterface interface {
	InsertNote(g *gin.Context)
	GetNote(g *gin.Context)
}
