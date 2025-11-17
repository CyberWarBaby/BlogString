package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", mainPage)
	server.GET("/blogs", getBlogs)
	server.POST("/blogs", createBlogs)
	server.GET("/blogs/:id", getBlog)
	server.PUT("/blogs/:id", updateBLog)
}
