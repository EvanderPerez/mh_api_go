package controllers

import (
	models "mh_api_go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers returns a list of users
func GetUsers(c *gin.Context) {
	users := []models.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
	}
	c.JSON(http.StatusOK, users)
}
