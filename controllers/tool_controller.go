package controllers

import (
	models "mh_api_go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers returns a list of users
func GetTools(c *gin.Context) {
	users := []models.Tool{
		{ID: 1, Model: "HP1640", ToolType: "Taladro", Brand: "Makita", Location: "Bodega"},
	}
	c.JSON(http.StatusOK, users)
}
