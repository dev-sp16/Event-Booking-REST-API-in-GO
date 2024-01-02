package requests

import (
	"net/http"
	"strconv"
	"event-booking.com/rest-api/events"
	"github.com/gin-gonic/gin"
)

func getEvents( context *gin.Context ) {
	events, err := events.GetAllEvents()

	if err != nil {
		context.JSON( http.StatusInternalServerError, gin.H{ "message": "Could not fetch events." } )
		return
	}

	context.JSON( http.StatusOK, events )
}

func getEvent( context *gin.Context ) {
	eventId, err := strconv.ParseInt( context.Param("id"), 10, 64 )

	if err != nil {
		context.JSON( http.StatusBadRequest, gin.H{ "message": "Could not parse event id." } )
		return
	}

	event, err := events.GetEventByID( eventId )

	if err != nil {
		context.JSON( http.StatusInternalServerError, gin.H{ "message": "Could not fetch event." } )
		return
	}

	context.JSON( http.StatusOK, event )
}

func createEvent( context *gin.Context ) {
	var event events.Event
	err := context.ShouldBindJSON( &event )

	if err != nil {
		context.JSON( http.StatusBadRequest, gin.H{ "message": "Could not parse request data." } )
		return
	}

	userID := context.GetInt64( "userId" )
	event.UserID = userID

	event.Save()

	if err != nil {
		context.JSON( http.StatusInternalServerError, gin.H{ "message": "Could not create event." } )
		return
	}

	context.JSON( http.StatusCreated, gin.H{ "message": "Event created!", "event": event } )
}

func updateEvent( context *gin.Context ) {
	eventId, err := strconv.ParseInt( context.Param("id"), 10, 64 )

	if err != nil {
		context.JSON( http.StatusBadRequest, gin.H{ "message": "Could not parse event id." } )
		return
	}

	userID := context.GetInt64( "userId" )
	event, err := events.GetEventByID( eventId )

	if err != nil {
		context.JSON( http.StatusInternalServerError, gin.H{ "message": "Could not fetch event." } )
		return
	}

	if event.UserID != userID {
		context.JSON( http.StatusUnauthorized, gin.H{ "message": "User is not authorized to update an event." } )
		return
	}

	var updatedEvent events.Event
	err = context.ShouldBindJSON( &updatedEvent )

	if err != nil {
		context.JSON( http.StatusBadRequest, gin.H{ "message": "Could not parse request data." } )
		return
	}

	updatedEvent.ID = eventId

	err = updatedEvent.Update()
	if err != nil {
		context.JSON( http.StatusInternalServerError, gin.H{ "message": "Could not update event." } )
		return
	}

	context.JSON( http.StatusOK, gin.H{ "message": "Event updated!" } )
}

func deleteEvent( context *gin.Context ) {
	eventId, err := strconv.ParseInt( context.Param( "id" ), 10, 64 ) // decimal int64

	if err != nil {
		context.JSON( http.StatusBadRequest, "Could not parse event id." )
		return
	}

	userID := context.GetInt64( "userId" )
	event, err := events.GetEventByID( eventId )

	if err != nil {
		context.JSON( http.StatusInternalServerError, gin.H{ "message": "Could not fetch event." } )
		return
	}

	if event.UserID != userID {
		context.JSON( http.StatusUnauthorized, gin.H{ "message": "User is not authorized to delete an event." } )
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON( http.StatusBadRequest, "Could not delete event." )
		return
	}

	context.JSON( http.StatusOK, gin.H{ "message": "Event deleted!" } )
}