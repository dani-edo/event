package models

import (
	"time"

	"edo.com/event/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"date_time"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, date_time, user_id)
	VALUES (?, ?, ?, ?, ?)` // ?: placeholder: prevent SQL injection
	stmt, err := db.DB.Prepare(query) // optional but good to use it
	if err != nil {
		return err
	}
	defer stmt.Close() // close the statement when we're done with it

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID) // Exec: insert, update, etc
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query) // Query: get data
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() { // Next: get the next row, return true if there is a next row
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}
