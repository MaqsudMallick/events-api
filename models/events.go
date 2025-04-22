package models

import (
	"log"
	"module-name/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

func (e *Event) SaveEvent() error {
	query := `INSERT INTO events (name, description, location, date_time, user_id) VALUES ($1, $2, $3, $4, $5)
	RETURNING id`
	err := db.DB.QueryRow(query, e.Name, e.Description, e.Location, e.DateTime, e.UserId).Scan(&e.ID)
	if err != nil {
		return err
	}
	log.Printf("Event with ID %d saved successfully", e.ID)
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events = []Event{}
	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (Event, error) {
	query := `SELECT * FROM events WHERE id = $1`
	var event Event
	err := db.DB.QueryRow(query, id).Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return event, err
	}
	return event, nil
}

func UpdateEventById(id int64, event Event) error {
	query := `UPDATE events SET name = $1, description = $2, location = $3, date_time = $4, user_id = $5 WHERE id = $6`
	_, err := db.DB.Exec(query, event.Name, event.Description, event.Location, event.DateTime, event.UserId, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEventById(id int64) error {
	query := `DELETE FROM events WHERE id = $1`
	_, err := db.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
