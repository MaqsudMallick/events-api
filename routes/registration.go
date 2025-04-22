package routes

import (
	"module-name/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerUserForEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var registration models.Registration
	registration.EventId = id
	registration.UserId = c.GetInt64("userId")
	err = registration.CreateRegistration()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Registration created successfully"})
}

func unregisterUserForEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var registration models.Registration
	registration.EventId = id
	registration.UserId = c.GetInt64("userId")
	err = registration.GetRegistrationId()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	err = models.DeleteRegistration(registration.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Registration deleted successfully"})
}
