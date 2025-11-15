package main

import (
	"fmt"

	"example.com/blog-api/db"
	"example.com/blog-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("This BlogString")
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
