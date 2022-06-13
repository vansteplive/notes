package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/vansteplive/notes-app-golang/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (p *AuthPostgres) CreateUser(user models.User) (int, error) {
	var idInt64 int64 = 2 << 32

	_, err := p.GetUser(user.Username, user.Password)
	if err == nil {
		return 0, errors.New("this username is already in use")
	}

	result, err := p.db.Exec("INSERT INTO users (username, password, firstname, lastname) VALUES ($1, $2, $3, $4)",
		user.Username, user.Password, user.Firstname, user.Lastname,
	)
	if err != nil {
		return 0, err
	}

	fmt.Printf("user data: %+v", user)

	idInt64, err = result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(idInt64), nil
}

func (p *AuthPostgres) GetUser(username, password string) (models.User, error) {
	var user models.User
	query := "SELECT id FROM users WHERE username=$1 AND password=$2"
	err := p.db.Get(&user, query, username, password)

	fmt.Printf("Selected user: %+v\n\ninput user: %+v", user, &models.User{Username: username, Password: password})

	return user, err
}
