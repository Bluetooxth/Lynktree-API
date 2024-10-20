package main

import (
	"os"

	"lynktree/config"
	"lynktree/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	config.ConnectDatabase()

	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://lynktree.vercel.app","https://lynktree.netlify.app","http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE","OPTIONS"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
	}))

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from lynktree Api",
		})
	})

	server.POST("/api/user/signup", controllers.Signup)
	server.POST("/api/user/login", controllers.Login)
	server.GET("/api/user/get-user/:username", controllers.GetUser)
	server.GET("/api/user/getuser-details", controllers.GetUserDetails)
	server.PUT("/api/user/update-user/:id", controllers.UpdateUser)
	server.DELETE("/api/user/delete-user/:id", controllers.DeleteUser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	server.Run(":" + port)
}
