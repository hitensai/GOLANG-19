package main

import (
	"example/reastapi/db"
	"example/reastapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
  db.Initdb()
  server := gin.Default()
  server.GET("/events", getEvents)
  server.POST("/events", createEvent)

  server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
  context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindBodyWithJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "coould not parse data"})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}