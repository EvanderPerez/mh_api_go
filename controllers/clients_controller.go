package controllers

import (
	"mh_api_go/database"
	models "mh_api_go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers returns a list of users
func GetClients(c *gin.Context) {
	var clients []models.Client

	result := database.DB.Find(&clients)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, clients)
}
