package requests

import (
	"net/http"

	"event-booking.com/rest-api/users"
	"event-booking.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup( context *gin.Context ) {
	var user users.User

	err := context.ShouldBindJSON( &user )

	if err != nil {
		context.JSON( http.StatusBadRequest, gin.H{ "message": "Could not parse request data." } )
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON( http.StatusInternalServerError, gin.H{ "message": "Could not save user." } )
		return
	}

	context.JSON( http.StatusCreated, gin.H{ "message": "User created!" } )
}

func login( context *gin.Context ) {
	var user users.User

	err := context.ShouldBindJSON( &user )

	if err != nil {
		context.JSON( http.StatusBadRequest, gin.H{ "message": "Could not parse request data." } )
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON( http.StatusUnauthorized, gin.H{ "message": err.Error() } )
		return
	}

	token, err := utils.GenerateToken( user.Email, user.ID )

	if err != nil {
		context.JSON( http.StatusInternalServerError, gin.H{ "message": err.Error() } )
	}

	context.JSON( http.StatusOK, gin.H{ "message": "Login successful!", "token": token } )
}