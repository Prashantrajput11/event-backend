package models // Declares this file belongs to the "models" package — usually used for data structures

import "time" // Imports Go's standard time package for handling date/time

// Event defines the structure of an event (like a class or a JS object shape)
type Event struct {
	ID          int       
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time  `binding:"required"`
	UserID      int       
}

// events is a slice (dynamic array) to store all event entries
// Think of it like: let events = [] in JavaScript
var events = []Event{}

// Save is a method on Event that adds the current event to the global events slice
// It's similar to doing events.push(event) in JavaScript
func (e Event) Save() {
	events = append(events, e) // Add the current Event (e) to the events list
}

// getAllEvents returns the list of all saved events
// Equivalent to: function GetAllEvents() { return events; } in JavaScript
func GetAllEvents() []Event {
	return events
}
