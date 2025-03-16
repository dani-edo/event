package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"date_time"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() {
	// later: add it to a database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
