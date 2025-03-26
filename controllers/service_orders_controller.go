package controllers

import (
	"mh_api_go/database"
	models "mh_api_go/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetUsers returns a list of users
func GetServiceOrders(c *gin.Context) {
	var service_orders []models.ServiceOrder
	clientIDStr := c.Query("client_id")

	var result *gorm.DB
	if clientIDStr != "" {
		result = database.DB.Where("clientID = ?", clientIDStr).Find(&service_orders)
	} else {
		result = database.DB.Find(&service_orders)
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, service_orders)
}
