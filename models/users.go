package models

import (
	"errors"

	"edo.com/event/db"
	"edo.com/event/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (user User) Save() error {
	query := `
		INSERT INTO users (email, password) VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hasedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hasedPassword)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	user.ID = id
	return err
}

func (user User) ValidateCredentials() error {
	query := `
		SELECT password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, user.Email)

	var retrievePassword string
	err := row.Scan(&retrievePassword)
	if err != nil {
		return errors.New("Invalid credentials")
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrievePassword)
	if !passwordIsValid {
		return errors.New("Invalid credentials")
	}
	return nil
}
