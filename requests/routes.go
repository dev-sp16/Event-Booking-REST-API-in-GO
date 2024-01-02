package requests

import (
	"event-booking.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET( "/events", getEvents )
	server.GET( "/events/:id", getEvent )

	authGroup := server.Group( "/" ) // group routes requiring middleware
	authGroup.Use( middlewares.Authenticate )
	authGroup.POST( "/events", createEvent )
	authGroup.PUT( "/events/:id", updateEvent )
	authGroup.DELETE( "/events/:id", deleteEvent )
	authGroup.POST( "/events/:id/register", registerForEvent )
	authGroup.DELETE( "/events/:id/register", cancelRegistration )

	server.POST( "/signup", signup )
	server.POST( "/login", login )
}