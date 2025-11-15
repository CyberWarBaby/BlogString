package routes

import (
	"net/http"

	"example.com/blog-api/models"
	"github.com/gin-gonic/gin"
)

func mainPage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "This BlogString"})
}

func getBlogs(context *gin.Context) {
	blogs, err := models.GetAllBlogs()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch blogs"})
	}
	context.JSON(http.StatusOK, blogs)
}
