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
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-ALlow-Methods", "POST,GET,PUT,DELETE,PATCH")

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

	server.Use(corsMiddleware())

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	server.Run(":" + port)
}