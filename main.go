package main

import (
	"gin_docker/controllers"
	"gin_docker/db"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// init gets called before the main function
func init() {
	// Log error if .env file does not exist
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

func main() {
	// Init gin router
	router := gin.Default()

	db := db.SetupModels() // new// Provide db variable to controllers
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	v1 := router.Group("/api/v1")
	{
		hello := new(controllers.HelloWorldController)

		v1.GET("/hello", hello.SampleHelloApi)
	}

	// Handle error response when a route is not defined
	router.NoRoute(func(c *gin.Context) {
		// In gin this is how you return a JSON response
		c.JSON(404, gin.H{"message": "404 Not found"})
	})

	// Init our server
	router.Run()
}
