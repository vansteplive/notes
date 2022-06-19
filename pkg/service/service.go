package service

import (
	"github.com/vansteplive/notes-app-golang/models"
	"github.com/vansteplive/notes-app-golang/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Note interface {
	CreateNote(userId int, note models.NoteItem) (int, error)
	GetAll(userId int) ([]models.NoteItem, error)
}

type Service struct {
	Authorization
	Note
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Note:          NewNoteItemPostgres(repos.NoteItem),
	}
}
