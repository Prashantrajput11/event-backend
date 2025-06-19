package main

import (
	"event-backend/models" // Importing models package containing the Event struct and functions
	"net/http"

	"github.com/gin-gonic/gin" // Gin web framework for routing and middleware
)

// main initializes the HTTP server, sets up routes, and starts the server on port 8080.
func main() {
	server := gin.Default() // Default Gin router with Logger and Recovery middleware

	server.GET("/events", getEvents)   // GET endpoint to retrieve all events
	server.POST("/events", createEvent) // POST endpoint to create a new event

	server.Run(":8080") // Start server on port 8080
}

// getEvents handles GET /events route.
// It retrieves all events from the model and returns them as a JSON array.
//
// Parameters:
// - context: *gin.Context — provides request and response context.
func getEvents(context *gin.Context) {
	events := models.GetAllEvents() // Retrieve all events
	context.JSON(http.StatusOK, events) // Respond with status 200 and the list of events
}

// createEvent handles POST /events route.
// It binds JSON payload to an Event struct, assigns mock IDs, and returns the created event.
//
// Parameters:
// - context: *gin.Context — provides request and response context.
func createEvent(context *gin.Context) {
	var event models.Event

	// Parse the incoming JSON body into the event struct
	err := context.ShouldBindJSON(&event)
	if err != nil {
		// Return error if parsing fails
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse"})
		return
	}

	// Temporary hardcoded values (should be replaced with real logic)
	event.ID = 1
	event.UserID = 1

	event.Save()

	// Return the created event with HTTP 201
	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   event,
	})
}
