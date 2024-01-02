package middlewares

import (
	"net/http"
	"event-booking.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate( context *gin.Context ) {
	token := context.Request.Header.Get( "Authorization" )

	if token == "" {
		context.AbortWithStatusJSON( http.StatusUnauthorized, gin.H{ "message": "Unauthorized user." } )
		return
	}

	userID, err := utils.VerifyToken( token )

	if err != nil {
		context.AbortWithStatusJSON( http.StatusUnauthorized, gin.H{ "message": "Unauthorized user." } )
		return
	}

	context.Set( "userId", userID )
	context.Next() // next handler should be executed
}