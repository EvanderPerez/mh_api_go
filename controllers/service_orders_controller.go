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

	context := database.DB
	context = context.Preload("Client").Preload("Tools")

	var result *gorm.DB
	if clientIDStr != "" {
		result = context.Where("client_ID = ?", clientIDStr).Find(&service_orders)
	} else {
		result = context.Find(&service_orders)
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, service_orders)
}

func CreateServiceOrder(c *gin.Context) {
	var order models.ServiceOrder

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&order) // Auto-creates Client, ServiceOrder & Tools
	c.JSON(http.StatusCreated, order)
}
