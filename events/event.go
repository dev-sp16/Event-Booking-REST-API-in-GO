package events

import (
	"time"
	"event-booking.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

var events = []Event{};

func ( event *Event ) Save() error {
	insertQuery := `INSERT INTO events( name, description, location, datetime, user_id)
	VALUES ( ?, ?, ?, ?, ? )
	`
	stmt, err := db.DB.Prepare( insertQuery )

	if err != nil {
		return err
	}

	defer stmt.Close() // close the statement
	result, err := stmt.Exec( event.Name, event.Description, event.Location, event.DateTime, event.UserID )

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	event.ID = id

	if err != nil {
		return err
	}

	return nil
}

func GetAllEvents() ( []Event, error )  {
	getQuery := "SELECT * FROM events"
	
	rows, err := db.DB.Query( getQuery ) // Exec generally is used in case of update/insert
	
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan( &event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID )

		if err != nil {
			return nil, err
		}

		events = append( events, event )
	}

	return events, nil
}

func GetEventByID( id int64 ) ( *Event, error ) {
	getQuery := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow( getQuery, id )

	var event Event
	err := row.Scan( &event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID )

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func ( event *Event ) Update() error {
	updateQuery := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare( updateQuery )

	if err != nil {
		return err
	}

	defer stmt.Close()

	_,err = stmt.Exec( event.Name, event.Description, event.Location, event.DateTime, event.ID )
	return err
}

func ( event Event ) Delete() error {
	deleteQuery := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare( deleteQuery )

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec( event.ID )
	return err
}

func ( event Event ) Register( userId int64 ) error {
	query := "INSERT INTO registrations( event_id, user_id ) VALUES ( ?, ? )"
	stmt, err := db.DB.Prepare( query )

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec( event.ID, userId )

	return err
}

func ( event Event ) CancelRegistration( userId int64 ) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"

	stmt, err := db.DB.Prepare( query )

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec( event.ID, userId )

	return err
}