package repositories

import (
	"database/sql"
	"list-notes/app/models"
	"log"
)

type NoteRepository struct {
	DB *sql.DB
}

func NewNoteRepository(db *sql.DB) NoteRepositoryInterface {
	return &NoteRepository{DB: db}
}

// InsertNote implements NoteRepositoryInterface
func (n *NoteRepository) InsertNote(post models.PostNote) bool {
	stmt, err := n.DB.Prepare("INSERT INTO notes (title, description) VALUES ($1, $2)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(post.Title, post.Description)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

// SelectNote implements NoteRepositoryInterface
func (n *NoteRepository) SelectNote() []models.Note {
	var result []models.Note
	rows, err := n.DB.Query("SELECT * FROM notes")
	if err != nil {
		log.Println(err)
		return nil
	}
	for rows.Next() {
		var (
			id          uint
			title       string
			description string
		)
		err := rows.Scan(&id, &title, &description)
		if err != nil {
			log.Println(err)
		} else {
			note := models.Note{Id: id, Title: title, Description: description}
			result = append(result, note)
		}
	}
	return result
}
