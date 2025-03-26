package routes

import (
	controllers "mh_api_go/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/users", controllers.GetUsers)
	router.GET("/tools", controllers.GetTools)
	router.GET("/service_orders", controllers.GetServiceOrders)
	router.GET("/clients", controllers.GetClients)

	return router
}
