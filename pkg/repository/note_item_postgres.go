package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/vansteplive/notes-app-golang/models"
)

type NoteItemPostgres struct {
	db *sqlx.DB
}

func NewNoteItemPostgres(db *sqlx.DB) *NoteItemPostgres {
	return &NoteItemPostgres{
		db: db,
	}
}

func (p *NoteItemPostgres) Create(userId int, item models.NoteItem) (int, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return 0, nil
	}
	var lastInsertId int

	createQuery := "INSERT INTO notes (title, description) VALUES ($1, $2) RETURNING id"
	if err = p.db.QueryRow(createQuery, item.Title, item.Description).Scan(&lastInsertId); err != nil {
		return 0, err
	}

	createUserNoteQuery := "INSERT INTO usersNotes (user_id, note_id) values ($1, $2)"
	if _, err := p.db.Exec(createUserNoteQuery, userId, lastInsertId); err != nil {
		tx.Rollback()
		return 0, err
	}

	return lastInsertId, tx.Commit()
}

func (p *NoteItemPostgres) GetAll(userId int) ([]models.NoteItem, error) {
	var list []models.NoteItem

	query := "SELECT nt.id, nt.title, nt.description FROM notes AS nt INNER JOIN usersNotes AS un ON nt.id = un.note_id WHERE un.user_id = $1"

	if err := p.db.Select(&list, query, userId); err != nil {
		return nil, err
	}

	return list, nil
}

// TODO: Реализовать функцию для получения записи по id
