package routes

import (
	"example.com/blog-api/models"
	"github.com/gin-gonic/gin"
)

func getBlogs(context *gin.Context) {
	blogs, err := models.GetAllBlogs
}
