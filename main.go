package main

import (
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

	// v1 := router.Group("/api/v1")
	// {
	// 	hello := new(controllers.HelloWorldController)

	// 	v1.GET("/hello", hello.sampleHelloAPI)
	// }

	// Handle error response when a route is not defined
	router.NoRoute(func(c *gin.Context) {
		// In gin this is how you return a JSON response
		c.JSON(404, gin.H{"message": "Not fodund"})
	})

	// Init our server
	router.Run()
}
