package main

import (
	"net/http"

	"github.com/JFMajer/rest-api-gin/db"
	"github.com/JFMajer/rest-api-gin/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context)  {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context)  {
	var event models.Event
	 err := context.ShouldBindJSON(&event)
	 if err != nil {
		 context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	 }
	 
	 _, err = event.Save()
	 if err != nil {
		 context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		 return
	 }
	 context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})


}