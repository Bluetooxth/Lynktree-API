package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"lynktree/config"
	"lynktree/controllers"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	_ = godotenv.Load()

	config.ConnectDatabase()

	server := gin.Default()

	server.Use(corsMiddleware())

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from lynktree Api",
		})
	})

	server.POST("/api/user/signup", controllers.Signup)
	server.POST("/api/user/login", controllers.Login)
	server.GET("/api/user/get-user/:username", controllers.GetUser)
	server.PUT("/api/user/update-user/:id", controllers.UpdateUser)
	server.DELETE("/api/user/delete-user/:id", controllers.DeleteUser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	server.Run(":" + port)
}