package main

import (
	"fmt"
	"net/http"

	"example.com/blog-api/db"
	"example.com/blog-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("This BlogString")
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)
	// global 404
	server.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "route not found",
		})
	})

	server.Run(":8080")
}
