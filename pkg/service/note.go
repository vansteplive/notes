package service

import (
	"github.com/vansteplive/notes-app-golang/models"
	"github.com/vansteplive/notes-app-golang/pkg/repository"
)

type NoteItemService struct {
	repos repository.NoteItem
}

func NewNoteItemPostgres(repos repository.NoteItem) *NoteItemService {
	return &NoteItemService{
		repos: repos,
	}
}

func (s *NoteItemService) CreateNote(item models.NoteItem) (int, error) {
	return s.repos.Create(item)
}

func (s *NoteItemService) GetAll() ([]models.NoteItem, error) {
	return s.repos.GetAll()
}
