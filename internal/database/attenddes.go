package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type AttenddesModel struct {
	DB *sql.DB
}

type Attenddes struct {
	Id      int `json:"id"`
	UserId  int `json:"userId"`
	EventId int `json:"eventId"`
}

func (m *AttenddesModel) Insert(attendee *Attenddes) (*Attenddes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// CHANGED: Use ? instead of $1, $2 for SQLite
	query := "INSERT INTO attendees (event_id, user_id) VALUES (?, ?)"

	result, err := m.DB.ExecContext(ctx, query, attendee.EventId, attendee.UserId)
	if err != nil {
		return nil, err
	}

	// CHANGED: Get the ID using LastInsertId for SQLite
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	attendee.Id = int(id)
	return attendee, nil
}

// func (m *AttenddesModel) GetByEventAndAttendees(eventId, userId int) (*Attenddes, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()
	
// 	// CHANGED: Use ? instead of $1, $2 for SQLite
// 	query := "SELECT id, user_id, event_id FROM attendees WHERE event_id = ? AND user_id = ?"
// 	var attendee Attenddes
	
// 	err := m.DB.QueryRowContext(ctx, query, eventId, userId).Scan(
// 		&attendee.Id, 
// 		&attendee.UserId, 
// 		&attendee.EventId,
// 	)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}
// 	return &attendee, nil
// }

func (m *AttenddesModel) GetByEventAndAttendees(eventId, userId int) (*Attenddes, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    
    query := "SELECT id, user_id, event_id FROM attendees WHERE event_id = ? AND user_id = ?"
    var attendee Attenddes
    
    // ADD DEBUG PRINT
    fmt.Printf("Executing query: %s with eventId: %d, userId: %d\n", query, eventId, userId)
    
    err := m.DB.QueryRowContext(ctx, query, eventId, userId).Scan(
        &attendee.Id, 
        &attendee.UserId, 
        &attendee.EventId,
    )
    if err != nil {
        // ADD DETAILED ERROR PRINTING
        fmt.Printf("Database error: %v\n", err)
        fmt.Printf("Error type: %T\n", err)
        
        if err == sql.ErrNoRows {
            fmt.Println("No rows found - attendee doesn't exist")
            return nil, nil
        }
        return nil, err
    }
    fmt.Printf("Attendee found: %+v\n", attendee)
    return &attendee, nil
}

func (m *AttenddesModel) GetAttenddesByEvent(eventId int) ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	// CHANGED: Use ? instead of $1 for SQLite
	query := `
	SELECT u.id, u.name, u.email
	FROM users u
	JOIN attendees a ON u.id = a.user_id
	WHERE a.event_id = ?
	`

	rows, err := m.DB.QueryContext(ctx, query, eventId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (m *AttenddesModel) Delete(userId, eventId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	// CHANGED: Use ? instead of $1, $2 for SQLite
	query := "DELETE FROM attendees WHERE user_id = ? AND event_id = ?"
	_, err := m.DB.ExecContext(ctx, query, userId, eventId)
	if err != nil {
		return err
	}
	return nil
}

func (m *AttenddesModel) GetEventsByAttendees(attendeeId int) ([]*Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	// CHANGED: Use ? instead of $1 for SQLite
	query := `
		SELECT e.id, e.owner_id, e.name, e.description, e.date, e.location
		FROM events e
		JOIN attendees a ON e.id = a.event_id
		WHERE a.user_id = ?
	`
	rows, err := m.DB.QueryContext(ctx, query, attendeeId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var events []*Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.OwnerID, &event.Name, &event.Description, &event.Date, &event.Location)
		if err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	return events, nil
}