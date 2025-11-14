package main

import (
	"fmt"

	"example.com/blog-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("This BlogString")

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
