package database

import (
	"context"
	"database/sql"
	"time"
)

type AttendeeModel struct {
	DB *sql.DB
}

type Attendee struct {
	Id      int `json:"id"`
	UserId  int `json:"userId"`
	EventId int `json:"eventId"`
}

func (m *AttendeeModel) Insert(attendee *Attendee) (*Attendee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "INSERT INTO attendees (user_id, event_id) VALUES ($1, $2) RETURNING id"

	// Get the new attendee ID
	err := m.DB.QueryRowContext(ctx, query, attendee.UserId, attendee.EventId).Scan(&attendee.Id)
	if err != nil {
		return nil, err
	}

	return attendee, nil
}

func (m *AttendeeModel) GetByEventAndAttentee(eventId, userId int) (*Attendee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT id, user_id, event_id FROM attendees WHERE event_id = $1 AND user_id = $2"

	var attendee Attendee
	err := m.DB.QueryRowContext(ctx, query, eventId, userId).Scan(&attendee.Id, &attendee.UserId, &attendee.EventId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No attendee found
		}
		return nil, err
	}

	return &attendee, nil
}

func (m *AttendeeModel) GetAttendeesByEvent(eventId int) ([]User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT u.id, u.name, u.email
		FROM users u
		JOIN attendees a ON u.id = a.user_id
		WHERE a.event_id = $1
	`

	rows, err := m.DB.QueryContext(ctx, query, eventId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Scan the results into a slice of User structs
	var attendees []User

	// Iterate over the rows
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		attendees = append(attendees, user)
	}

	// Check for any errors encountered during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return attendees, nil
}

func (m *AttendeeModel) Delete(userId, eventId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "DELETE FROM attendees WHERE user_id = $1 AND event_id = $2"

	_, err := m.DB.ExecContext(ctx, query, userId, eventId)
	if err != nil {
		return err
	}
	return nil
}

func (m *AttendeeModel) GetByAttendee(attendeeId int) ([]Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT e.id, e.owner_id, e.name, e.description, e.date, e.location
		FROM events e
		JOIN attendees a ON e.id = a.event_id
		WHERE a.user_id = $1
	`
	rows, err := m.DB.QueryContext(ctx, query, attendeeId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Scan the results into a slice of Event structs
	var events []Event
	// Iterate over the rows
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.OwnerId, &event.Name, &event.Description, &event.Date, &event.Location)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	// Check for any errors encountered during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
