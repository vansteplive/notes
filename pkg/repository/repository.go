package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/vansteplive/notes-app-golang/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type NoteItem interface {
	Create(userId int, item models.NoteItem) (int, error)
	GetAll(userId int) ([]models.NoteItem, error)
}

type Repository struct {
	Authorization
	NoteItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		NoteItem:      NewNoteItemPostgres(db),
	}
}
