package routes

import (
	"net/http"
	"strconv"

	"edo.com/event/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get events", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event) // bind the request body to the event struct

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create event", "error": err.Error()})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID", "error": err.Error()})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, *event)
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID", "error": err.Error()})
		return
	}

	_, err = models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event", "error": err.Error()})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent) // bind the request body to the event struct

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create event", "error": err.Error()})
		return
	}

	updatedEvent.ID = id
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated", "event": updatedEvent})
}
