package routes

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/JFMajer/rest-api-gin/models"
	"github.com/gin-gonic/gin"
)

func getRegistrations(context *gin.Context) {
	registrations, err := models.GetRegistrations()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, registrations)
}

func registerForEvent(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEvent(id)
	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
			return
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	err = event.RegisterUser(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "success"})
}

func cancelRegistration(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEvent(id)
	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
			return
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
