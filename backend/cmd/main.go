package main

import (
	"AUTH/src/database"
	"AUTH/src/dropbox"
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
		AllowOrigins: 	 []string{"http://localhost:3000"},
		AllowMethods:    []string{"OPTIONS", "GET", "PUT", "POST", "DELETE", "PATCH"},
		AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:          time.Hour * 24,
	}))

	api := router.Group("/api")
	{
		api.POST("/login", handlers.Login)
		api.GET("/auth/callback", handlers.HandleDropboxCallback)
		api.GET("/dropbox_login", handlers.HandleDropboxLogin)
		api.GET("/folders", func(c *gin.Context) {
			path := c.Query("path") 
			dropbox.GetDropboxFolders(c, path) 
		})
		api.POST("/register", handlers.Register)
		secured := api.Group("/secured").Use(middleware.AuthenticationMiddleware())
		{
			secured.GET("/test", handlers.Test)
			secured.GET("/me", func(ctx *gin.Context) { ctx.Status(200) })
		}
	}
	return router
}
