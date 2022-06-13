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

func (p *NoteItemPostgres) Create(item models.NoteItem) (int, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return 0, nil
	}

	var idInt64 int64 = 2 << 32
	createQuery := "INSERT INTO notes (title, description) VALUES ($1, $2)"
	row, err := p.db.Exec(createQuery, item.Title, item.Description)
	if err != nil {
		return 0, nil
	}

	idInt64, err = row.LastInsertId()
	if err != nil {
		return 0, err
	}
	itemId := int(idInt64)

	return itemId, tx.Commit()
}

func (p *NoteItemPostgres) GetAll() ([]models.NoteItem, error) {
	var noteList []models.NoteItem

	query := "SELECT * FROM notes"

	if err := p.db.Select(&noteList, query); err != nil {
		return nil, err
	}

	return noteList, nil
}

// TODO: Реализовать функцию для получения записи по id
