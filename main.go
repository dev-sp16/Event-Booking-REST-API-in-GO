package main

import (
	"event-booking.com/rest-api/requests"
	"event-booking.com/rest-api/db"
	"github.com/gin-gonic/gin" // gin framework for handling REST APIs
)

func main() {
	db.InitDB()

	server := gin.Default();

	requests.RegisterRoutes( server )

	server.Run( ":8080" ) // localhost:8080
}
