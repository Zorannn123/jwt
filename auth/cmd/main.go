package main

import (
	"AUTH/src/database"
	"AUTH/src/handlers"
	"AUTH/src/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.Migrate()
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/login", handlers.Login)
		api.POST("/register", handlers.Register)
		secured := api.Group("/secured").Use(middleware.AuthenticationMiddleware())
		{
			secured.GET("/test", handlers.Test)
		}
	}
	return router
}
