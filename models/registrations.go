package models

import "module-name/db"

type Registration struct {
	ID      int64
	EventId int64 `binding:"required"`
	UserId  int64
}

func (r *Registration) CreateRegistration() error {
	query := `INSERT INTO registrations (event_id, user_id) VALUES ($1, $2)`
	_, err := db.DB.Exec(query, r.EventId, r.UserId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRegistration(id int64) error {
	query := `DELETE FROM registrations WHERE id = $1`
	_, err := db.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Registration) GetRegistrationId() error {
	query := `SELECT id FROM registrations WHERE event_id = $1 AND user_id = $2`
	err := db.DB.QueryRow(query, r.EventId, r.UserId).Scan(&r.ID)
	if err != nil {
		return err
	}
	return nil
}
