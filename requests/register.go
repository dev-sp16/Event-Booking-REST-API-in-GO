package requests

import (
	"net/http"
	"strconv"

	"event-booking.com/rest-api/events"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context ) {
	userId := context.GetInt64( "userId" )
	eventId, err := strconv.ParseInt( context.Param( "id" ), 10, 64 ) // decimal int64

	if err != nil {
		context.JSON( http.StatusBadRequest, "Could not parse event id." )
		return
	}

	event, err := events.GetEventByID( eventId )

	if err != nil {
		context.JSON( http.StatusInternalServerError, gin.H{ "message": "Could not fetch event." } )
		return
	}

	err = event.Register( userId )

	if err != nil {
		context.JSON( http.StatusInternalServerError, gin.H{ "message": "Could notregister user for event." } )
		return
	}

	context.JSON( http.StatusCreated, gin.H{ "message": "Registered succesfully!" } )
}

func cancelRegistration(context *gin.Context ) {
	userId := context.GetInt64( "userId" )
	eventId, err := strconv.ParseInt( context.Param( "id" ), 10, 64 ) // decimal int64

	if err != nil {
		context.JSON( http.StatusBadRequest, "Could not parse event id." )
		return
	}

	var event events.Event
	event.ID = eventId

	err = event.CancelRegistration( userId )

	if err != nil {
		context.JSON( http.StatusInternalServerError, gin.H{ "message": "Could not cancel registration." } )
		return
	}

	context.JSON( http.StatusOK, gin.H{ "message": "Registered cancelled succesfully!" } )
}