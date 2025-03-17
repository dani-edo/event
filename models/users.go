package models

import "edo.com/event/db"

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
		panic(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Email, user.Password)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	user.ID = id
	return err
}
