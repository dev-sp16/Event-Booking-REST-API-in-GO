package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ( string, error ) {
	hashedPass, err := bcrypt.GenerateFromPassword( []byte( password ), 14 ) // hashing cost = 14
	return string( hashedPass ), err
}

func ComparePasswords( password, hashedPassword string ) bool {
	err := bcrypt.CompareHashAndPassword( []byte( hashedPassword ), []byte( password ) )
	return err == nil
}