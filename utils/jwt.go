package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "superset"

func GenerateToken( email string, userId int64 ) ( string, error ) {
	token := jwt.NewWithClaims( jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add( time.Hour * 2 ).Unix(), // 2 hours
	} )

	return token.SignedString( []byte( secretKey ) ) // send as byte slice
}

func VerifyToken( token string ) ( int64, error ) {
	parsedToken, err := jwt.Parse( token, func( token *jwt.Token ) ( interface{}, error ) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // type checking 

		if !ok {
			return nil, errors.New( "Unexpected signing method." )
		}
		return []byte( secretKey ), nil
	} )

	if err != nil {
		return 0, errors.New( "Could not parse token." )
	}

	isValid := parsedToken.Valid

	if !isValid {
		return 0, errors.New( "Invalid token!" )
	}

	claims, ok := parsedToken.Claims.( jwt.MapClaims )

	if !ok {
		return 0, errors.New( "Invalid token claims." )
	}

//	email := claims[ "email" ].( string )
	userId := int64( claims[ "userId"].( float64 ) ) // value stored as float64 internally

	return userId, nil
}