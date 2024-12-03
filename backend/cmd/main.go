package main

import (
	"AUTH/src/database"
	"AUTH/src/handlers"
	"AUTH/src/middleware"
	"time"

	"github.com/gin-contrib/cors"
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

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"OPTIONS", "GET", "PUT", "POST", "DELETE", "PATCH"},
		AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		MaxAge:          time.Hour * 24,
	}))

	api := router.Group("/api")
	{
		api.POST("/login", handlers.Login)
		api.POST("/register", handlers.Register)
		secured := api.Group("/secured").Use(middleware.AuthenticationMiddleware())
		{
			secured.GET("/test", handlers.Test)
			secured.GET("/me", func(ctx *gin.Context) { ctx.Status(200) })
		}
	}
	return router
}
