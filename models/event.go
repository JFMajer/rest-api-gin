package models

import (
	"time"

	"github.com/JFMajer/rest-api-gin/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

type Registration struct {
	ID      int64
	UserID  int64
	EventID int64
}

func (e *Event) Save() (int, error) {
	query :=
		`INSERT INTO events (name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	result, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return 0, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	e.ID = int64(lastId)
	return int(lastId), err
}

func GetEvent(id int) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	event := &Event{}
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func GetAllEvents() ([]*Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []*Event{}
	for rows.Next() {
		event := &Event{}
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func (e *Event) Update() error {
	query := "UPDATE events SET name = ?, description = ?, location = ?, dateTime = ? WHERE id = ?"
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	if err != nil {
		return err
	}
	return nil
}

func Delete(id int) error {
	query := "DELETE FROM events WHERE id = ?"
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func GetRegistrations() ([]*Registration, error) {
	query := "SELECT * FROM registrations"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	registrations := []*Registration{}
	for rows.Next() {
		registration := &Registration{}
		err := rows.Scan(&registration.ID, &registration.UserID, &registration.EventID)
		if err != nil {
			return nil, err
		}
		registrations = append(registrations, registration)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return registrations, nil

}

func (e *Event) RegisterUser(userId int64) error {
	query := "INSERT into registrations (user_id, event_id) VALUES (?, ?)"
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(userId, e.ID)
	if err != nil {
		return err
	}
	return nil
}

func (e *Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE user_id = ? AND event_id = ?"
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(userId, e.ID)
	if err != nil {
		return err
	}
	return nil
}
