package models

type User struct {
	ID        int    `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	Password  string `json:"password" db:"password"`
	Firstname string `json:"firstname" db:"firstname"`
	Lastname  string `json:"lastname" db:"lastname"`
}
