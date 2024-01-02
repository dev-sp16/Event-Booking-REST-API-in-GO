package users

import (
	"errors"

	"event-booking.com/rest-api/db"
	"event-booking.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	insertQuery := `
	INSERT INTO users( email, password )
	VALUES ( ?, ? )
	`
	stmt, err := db.DB.Prepare(insertQuery)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPass, err := utils.HashPassword( user.Password )

	if err != nil {
		return err
	}

	result, err := stmt.Exec( user.Email, hashedPass )

	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()

	user.ID = userID

	return err
}

func ( user *User ) ValidateCredentials() error {
	query := "SELECT id,password FROM users WHERE email = ?"

	row := db.DB.QueryRow( query, user.Email )

	var retrievedPass string
	err := row.Scan( &user.ID, &retrievedPass )

	if err != nil {
		return errors.New( "Invalid credentials." )
	}

	validPass := utils.ComparePasswords( user.Password, retrievedPass )

	if !validPass {
		return errors.New( "Invalid credentials." )
	}

	return nil
}