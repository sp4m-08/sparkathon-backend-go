package main

import (
	"log"
	"upc-backend-sparkathon/main/routes"
	"upc-backend-sparkathon/main/services"

	"github.com/gin-gonic/gin"
)

func main() {
	services.InitDatabase()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{"message": `server is running on port 8080!`})
	})
	routes.RegisterProductRoutes(r)

	port := "8080"
	log.Printf("Server running at: http://localhost:%s\n", port)

	err := r.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
