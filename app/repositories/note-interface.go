package repositories

import "list-notes/app/models"

type NoteRepositoryInterface interface {
	SelectNote() []models.Note
	InsertNote(post models.PostNote) bool
}
