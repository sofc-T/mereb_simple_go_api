package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sofc-t/mereb_simple_go_api/internals/handlers"
)

func main() {
	router := gin.Default()
	
	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})	
	router.GET("/person", handlers.GetPersons)
	router.GET("/person/:id", handlers.GetPerson)
	router.POST("/person", handlers.CreatePerson)
	router.PUT("/person/:id", handlers.UpdatePerson)
	router.DELETE("/person/:id", handlers.DeletePerson)
	
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Endpoint not found"})
	})

	router.Run(":8080")
}
