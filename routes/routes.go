package routes

import (
	"fmt"
	controllers "mh_api_go/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// router.Use(cors.Default())
	router.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			fmt.Println("Incoming Origin:", origin)
		}
		c.Next()
	})

	api := router.Group("/api/v1")
	{
		api.GET("/users", controllers.GetUsers)
		api.GET("/tools", controllers.GetTools)
		api.GET("/service_orders", controllers.GetServiceOrders)
		api.GET("/clients", controllers.GetClients)
		api.GET("/clients/names", controllers.GetClientNames)
		api.POST("/service_order", controllers.CreateServiceOrder)

	}

	return router
}
