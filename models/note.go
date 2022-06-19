package models

type UserNotes struct {
	ID     int `json:"id" db:"db"`
	UserID int `json:"user_id" db:"user_db"`
	NoteID int `json:"note_id" db:"note_db"`
}

type NoteItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}
